package store

import (
	"database/sql"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/pborman/uuid"
	"github.com/pkg/errors"
)

const (
	// codeDuplicateInsert is the code pg throws on a duplicate insert for a unique constraint
	codeDuplicateInsert = pq.ErrorCode("23505")
)

var (
	// ErrDuplicateStore represents an error when inserting a non unique store
	ErrDuplicateStore = errors.New("store location and name must be unique")
)

// Store represents an annoy store
type Store struct {
	ID       uuid.UUID `json:"id" db:"store_id"`
	Location string    `json:"-" db:"location"`
	Created  time.Time `json:"created" db:"created"`
	Updated  time.Time `json:"updated" db:"updated"`
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

// CreateStoreIfNotExists creates a new store if it does not exist or returns it if it does
func (serv *Service) CreateStoreIfNotExists(s *Store) (*Store, bool, error) {
	st, err := serv.GetStore()
	if err != nil {
		if err := errors.Cause(err); err != sql.ErrNoRows {
			return nil, false, errors.Wrap(err, "get store by name")
		}
	}

	// didn't find the store so we should create it
	if err := errors.Cause(err); err == sql.ErrNoRows {
		stmt, err := serv.DB.PrepareNamed(insertStore)
		if err != nil {
			return nil, false, errors.Wrap(err, "preparing insert store query")
		}
		defer stmt.Close()

		args := map[string]interface{}{
			"id":       uuid.New(),
			"location": s.Location,
		}

		var r Store
		if err := stmt.Get(&r, args); err != nil {
			if pgErr, ok := err.(*pq.Error); ok {
				if pgErr.Code == codeDuplicateInsert {
					return nil, false, ErrDuplicateStore
				}
			}

			return nil, false, errors.Wrap(err, "insert store query")
		}

		return &r, false, nil
	}

	// found the store return it
	return st, true, nil
}

// GetStore retrieves a store from the database
func (serv *Service) GetStore() (*Store, error) {
	var r Store
	if err := serv.DB.Get(&r, getStore); err != nil {
		return nil, errors.Wrap(err, "get store query")
	}

	return &r, nil
}
