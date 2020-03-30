package document

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
	"time"
	"unicode"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pborman/uuid"
	"github.com/pkg/errors"
)

const (
	// TypeMarkdown is a markdown document_type_id
	TypeMarkdown = iota + 1
	// TypePDF is a PDF document_type_id
	TypePDF
	// TypeUnsupported is a unsupported document_type_id
	TypeUnsupported

	// codeDuplicateInsert is the code pg throws on a duplicate insert for a unique constraint
	codeDuplicateInsert = pq.ErrorCode("23505")

	// reservedFolderName is a reserved name for the root folder
	reservedFolderName = "_rootfolder_"
)

var (
	// ErrDuplicateIndex is returned when an index is being inserted and it already exists
	ErrDuplicateIndex = errors.New("index already exists")

	// ErrDuplicateDocument is returned when a document is being inserted and it already exists
	ErrDuplicateDocument = errors.New("document already exists")
)

// Type represents a document type
type Type struct {
	ID      int       `json:"id" db:"document_type_id"`
	Name    string    `json:"name" db:"name"`
	Created time.Time `json:"created" db:"created"`
	Updated time.Time `json:"updated" db:"updated"`
}

// SearchResult represents search results
type SearchResult struct {
	DocumentID   uuid.UUID `json:"id" db:"document_id"`
	SentenceID   uuid.UUID `json:"sentenceID" db:"sentence_id"`
	AnnoyID      int       `json:"annoyID" db:"annoy_id"`
	DocumentName string    `json:"name" db:"name"`
	Path         string    `json:"path" db:"path"`
	DownloadURL  string    `json:"downloadURL" db:"download_url"`
	Text         string    `json:"text" db:"sentence_text"`
	Rank         float32   `json:"rel" db:"rel"`
}

// Folder represents a folder
type Folder struct {
	FolderID uuid.UUID `json:"id" db:"folder_id"`
	Name     string    `json:"name" db:"name"`
	Path     string    `json:"path" db:"path"`
	Created  time.Time `json:"created" db:"created"`
	Updated  time.Time `json:"updated" db:"updated"`
}

// Document represents a document that was indexed
type Document struct {
	ID               uuid.UUID `json:"id" db:"document_id"`
	FolderID         uuid.UUID `json:"folderID" db:"folder_id"`
	TypeID           int       `json:"typeID" db:"document_type_id"`
	ObjectStorageURL string    `json:"objectStorageURL" db:"object_storage_url"`
	DownloadURL      string    `json:"downloadURL" db:"download_url"`
	Path             string    `json:"path" db:"path"`
	Name             string    `json:"name" db:"name"`
	Body             string    `json:"body" db:"body"`
	Created          time.Time `json:"created" db:"created"`
	Updated          time.Time `json:"updated" db:"updated"`
}

// Sentence represents indexed sentence from a document
type Sentence struct {
	ID         uuid.UUID       `db:"sentence_id"`
	DocumentID uuid.UUID       `db:"document_id"`
	StoreID    uuid.UUID       `db:"store_id"`
	AnnoyID    int             `db:"annoy_id"`
	Embedding  json.RawMessage `db:"embedding"`
	Body       string          `db:"body"`
	Context    string          `db:"context"`
	Created    time.Time       `json:"created" db:"created"`
	Updated    time.Time       `json:"updated" db:"updated"`
}

// GetEmbeddings returns the json value of embedding to a []float32
func (s *Sentence) GetEmbeddings() ([]float32, error) {
	var embeddings []float32
	if err := json.Unmarshal(s.Embedding, &embeddings); err != nil {
		return nil, errors.Wrap(err, "unmarshalling embedding")
	}

	return embeddings, nil
}

// Service contains functionality for managing the document data service
type Service struct {
	DB *sqlx.DB
}

// NewService returns a new agent service
func NewService(db *sqlx.DB) *Service {
	s := Service{
		DB: db,
	}

	return &s
}

// GetIndexContent gets all content previously indexed to rebuild the index
func (s *Service) GetIndexContent() ([]Sentence, error) {
	var r []Sentence
	if err := s.DB.Select(&r, getIndexContent); err != nil {
		return nil, errors.Wrap(err, "get index content query")
	}

	return r, nil
}

// cleanPath removes all whitespace and hyphens from ltree paths
func cleanPath(path string) string {
	path = strings.Replace(path, ".", "_", -1) // replace '.' with '_'
	path = strings.Replace(path, "/", ".", -1) // replace '/' with path delimeter '.'

	cleaned := strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) || r == '-' {
			return -1
		}

		return r
	}, path)

	return cleaned
}

// GetFolderByPath gets a folder given a specific path
func (s *Service) GetFolderByPath(path string) (*Folder, error) {
	stmt, err := s.DB.PrepareNamed(getFolderByPath)
	if err != nil {
		return nil, errors.Wrap(err, "preparing get folder by path query")
	}
	defer stmt.Close()

	args := map[string]interface{}{
		"path": cleanPath(path),
	}

	var r Folder
	if err := stmt.Get(&r, args); err != nil {
		return nil, errors.Wrap(err, "get folder by path query")
	}

	return &r, nil
}

// CreateFolderPath creates the folder a document is placed in if it does not exist.
// Parameter 'recursing' should be false when calling this method
func (s *Service) CreateFolderPath(path string, recursing bool) (string, error) {
	stmt, err := s.DB.PrepareNamed(insertFolder)
	if err != nil {
		return "", errors.Wrap(err, "preparing insert folder query")
	}
	defer stmt.Close()

	// get last "/" to decipher folder name
	slashIndex := strings.LastIndex(path, "/")

	// get folder name
	var name string
	if path == "/" {
		name = reservedFolderName
	}

	if (len(path) >= slashIndex+1) && name != reservedFolderName {
		name = path[slashIndex+1:]
	} else {
		name = reservedFolderName
	}

	// ensure that path is always set after root folder
	if path == "/" {
		path = reservedFolderName
	} else {
		if !recursing {
			path = fmt.Sprintf("%v/%v", reservedFolderName, path)
		}
	}

	folderID := uuid.New()
	args := map[string]interface{}{
		"id":   folderID,
		"name": name,
		"path": cleanPath(path),
	}

	// insert does not error on duplicates, check rows affected to check if we need to get the folder_id
	r, err := stmt.Exec(args)
	if err != nil {
		return "", errors.Wrap(err, "insert folder query")
	}

	a, err := r.RowsAffected()
	if err != nil {
		return "", errors.Wrap(err, "getting rows affected")
	}

	// get folder since insert returned nothing
	if a == 0 {
		f, err := s.GetFolderByPath(path)
		if err != nil {
			return "", errors.Wrap(err, "get folder by path")
		}

		folderID = f.FolderID.String()
		path = f.Path
	}

	// go up the path tree and create parent folders if we have a nested path
	parentPath := path

	// skip if we are at root folder path (reservedFolderName)
	if name != reservedFolderName {
		parentFolders := strings.Split(path, "/")

		for i := 0; i < len(parentFolders); i++ {
			// get one folder up
			slashIndex := strings.LastIndex(parentPath, "/")
			if slashIndex > -1 {
				parentPath = parentPath[:slashIndex]
			} else {
				parentPath = "/" // if we don't have anymore slashes set to root
			}

			_, err := s.CreateFolderPath(parentPath, true)
			if err != nil {
				return "", errors.Wrapf(err, "creating parent paths, err on path %v", parentPath)
			}
		}
	}

	return folderID, nil
}

// DeleteDocument deletes a document and all referencing data
func (s *Service) DeleteDocument(path string) error {
	stmt, err := s.DB.PrepareNamed(deleteDocumentByPath)
	if err != nil {
		return errors.Wrap(err, "preparing delete document by path query")
	}
	defer stmt.Close()

	args := map[string]interface{}{
		"path": cleanPath(path),
	}

	// sentence table has on delete cascade foreign key so all previous sentences will be removed
	if _, err := stmt.Exec(args); err != nil {
		return errors.Wrap(err, "error deleting document")
	}

	return nil
}

// GetDocumentByPath gets a document given a specific path
func (s *Service) GetDocumentByPath(path string) (*Document, error) {
	stmt, err := s.DB.PrepareNamed(getDocumentByPath)
	if err != nil {
		return nil, errors.Wrap(err, "preparing get document by path query")
	}
	defer stmt.Close()

	// ensure a file is set in the path
	filename := filepath.Base(path)
	if filename == "." {
		err := errors.New("a file was not defined in the given path")
		return nil, err
	}

	// ensure that all paths have a leading '/'
	if path[0] != '/' {
		path = "/" + path
	}

	// replace leading slash with root folder
	path = strings.Replace(path, "/", reservedFolderName, 1)

	// remove filename from path because we don't want to clean it
	path = strings.Replace(path, filename, "", 1)

	args := map[string]interface{}{
		"path": cleanPath(path) + "." + filename,
	}

	var r Document
	if err := stmt.Get(&r, args); err != nil {
		return nil, errors.Wrap(err, "get document by path query")
	}

	return &r, nil
}

// CreateDocument creates a new document
func (s *Service) CreateDocument(d *Document) (*Document, error) {
	stmt, err := s.DB.PrepareNamed(insertDocument)
	if err != nil {
		return nil, errors.Wrap(err, "preparing insert document query")
	}
	defer stmt.Close()

	// create folder given path, get folder_id and insert doc with it
	folderID, err := s.CreateFolderPath(d.Path, false)
	if err != nil {
		return nil, errors.Wrap(err, "creating folder for document")
	}

	// don't build sub dir if path is the root
	path := reservedFolderName
	if d.Path != "/" {
		path = fmt.Sprintf("%v/%v", reservedFolderName, d.Path)
	}

	args := map[string]interface{}{
		"id":               uuid.New(),
		"typeID":           d.TypeID,
		"name":             d.Name,
		"body":             d.Body,
		"path":             cleanPath(path) + "." + d.Name,
		"objectStorageURL": d.ObjectStorageURL,
		"downloadURL":      d.DownloadURL,
		"folderID":         folderID,
	}

	var r Document
	if err := stmt.Get(&r, args); err != nil {
		if pgErr, ok := err.(*pq.Error); ok {
			if pgErr.Code == codeDuplicateInsert {
				return nil, ErrDuplicateDocument
			}
		}

		return nil, errors.Wrap(err, "insert document query")
	}

	return &r, nil
}

// CreateSentence stores sentence and mapping to annoy index id from a parsed document
func (s *Service) CreateSentence(i *Sentence) (*Sentence, error) {
	stmt, err := s.DB.PrepareNamed(insertSentence)
	if err != nil {
		return nil, errors.Wrap(err, "preparing insert sentence query")
	}
	defer stmt.Close()

	args := map[string]interface{}{
		"id":         uuid.New(),
		"documentID": i.DocumentID,
		"storeID":    i.StoreID,
		"body":       i.Body,
		"embedding":  i.Embedding,
		"context":    i.Context,
	}

	var r Sentence
	if err := stmt.Get(&r, args); err != nil {
		return nil, errors.Wrap(err, "insert sentence query")
	}

	return &r, nil
}

// GetSearchResults returns documents given sentence ids. Spotify annoy returns a list of ids which relates to sentences
func (s *Service) GetSearchResults(sIDs []int) ([]SearchResult, error) {
	query, args, err := sqlx.In(getDocumentsBySentence, sIDs)
	if err != nil {
		return nil, errors.Wrap(err, "preparing get search results query")
	}

	query = s.DB.Rebind(query)

	var r []SearchResult
	if err := s.DB.Select(&r, query, args...); err != nil {
		return nil, errors.Wrap(err, "search results query")
	}

	return r, nil
}

// FullTextSearch performs postgres FTS on sentence bodies to get SearchResults
func (s *Service) FullTextSearch(text string) ([]SearchResult, error) {
	stmt, err := s.DB.PrepareNamed(fullTextSearchSentences)
	if err != nil {
		return nil, errors.Wrap(err, "preparing fts query")
	}
	defer stmt.Close()

	args := map[string]interface{}{
		"search": text,
	}

	var r []SearchResult
	if err := stmt.Select(&r, args); err != nil {
		return nil, errors.Wrap(err, "ftsquery")
	}

	return r, nil
}
