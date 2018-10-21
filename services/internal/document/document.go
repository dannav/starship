package document

import (
	"encoding/json"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pborman/uuid"
	"github.com/pkg/errors"
)

const (
	// TypeMarkdown is a markdown document_type_id
	TypeMarkdown = iota + 1
	// TypeText is a text document_type_id
	TypeText

	// codeDuplicateInsert is the code pg throws on a duplicate insert for a unique constraint
	codeDuplicateInsert = pq.ErrorCode("23505")
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
	AnnoyID      int       `json:"annoyID" db:"annoy_id"`
	DocumentName string    `json:"name" db:"name"`
	Text         string    `json:"text" db:"sentence_text"`
}

// Document represents a document that was indexed
type Document struct {
	ID      uuid.UUID `json:"id" db:"document_id"`
	TypeID  int       `json:"typeID" db:"document_type_id"`
	Name    string    `json:"name" db:"name"`
	TeamID  string    `json:"teamID" db:"team_id"` // TODO convert to uuid.UUID when auth is in place
	Body    string    `json:"body" db:"body"`
	Created time.Time `json:"created" db:"created"`
	Updated time.Time `json:"updated" db:"updated"`
}

// Sentence represents indexed sentence from a document
type Sentence struct {
	ID         uuid.UUID       `db:"sentence_id"`
	DocumentID uuid.UUID       `db:"document_id"`
	StoreID    uuid.UUID       `db:"store_id"`
	AnnoyID    int             `db:"annoy_id"`
	Embedding  json.RawMessage `db:"embedding"`
	Body       string          `db:"body"`
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

// GetIndexContentForTeam gets all content previously indexed for a team for rebuilding the index
func (s *Service) GetIndexContentForTeam(teamID string) ([]Sentence, error) {
	stmt, err := s.DB.PrepareNamed(getIndexContentForTeam)
	if err != nil {
		return nil, errors.Wrap(err, "preparing get index content query")
	}
	defer stmt.Close()

	args := map[string]interface{}{
		"teamID": teamID,
	}

	var r []Sentence
	if err := stmt.Select(&r, args); err != nil {
		return nil, errors.Wrap(err, "get index content query")
	}

	return r, nil
}

// CreateDocument creates a new document
func (s *Service) CreateDocument(d *Document) (*Document, error) {
	stmt, err := s.DB.PrepareNamed(insertDocument)
	if err != nil {
		return nil, errors.Wrap(err, "preparing insert document query")
	}
	defer stmt.Close()

	args := map[string]interface{}{
		"id":     uuid.New(),
		"typeID": d.TypeID,
		"teamID": d.TeamID,
		"name":   d.Name,
		"body":   d.Body,
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
