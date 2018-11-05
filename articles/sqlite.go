package articles

import (
	"database/sql"
	"errors"
)

// ErrSQLiteEmptyID is the error used when an empty string is given as an article is.
var ErrSQLiteEmptyID = errors.New("articles.SQLiteStore: Can't use empty string as id")

// SQLiteStore is a SQLite ArticleStore.
type SQLiteStore struct {
	addStmt    *sql.Stmt
	getStmt    *sql.Stmt
	updateStmt *sql.Stmt
	removeStmt *sql.Stmt
}

var _ Store = &SQLiteStore{}

// Add an article [a] to the SQLite database with the ID [id].
func (s *SQLiteStore) Add(id string, a Article) error {
	if id == "" {
		return ErrSQLiteEmptyID
	}
	_, err := s.addStmt.Exec(id, a.Title, a.Details, a.Author, a.Content)
	return err
}

// Get an article with the id [id] from the SQLite database.
func (s *SQLiteStore) Get(id string) (Article, error) {
	a := Article{}
	if id == "" {
		return a, ErrSQLiteEmptyID
	}
	err := s.getStmt.QueryRow(id).Scan(&a.Title, &a.Details, &a.Author, &a.Content)
	return a, err
}

// Update updates the article with id [id] with the data from [a].
func (s *SQLiteStore) Update(id string, a Article) error {
	if id == "" {
		return ErrSQLiteEmptyID
	}
	_, err := s.updateStmt.Exec(a.Title, a.Details, a.Author, a.Content, id)
	return err
}

// Remove an article with the id [id] from the SQLite database.
func (s *SQLiteStore) Remove(id string) error {
	if id == "" {
		return nil
	}
	_, err := s.removeStmt.Exec(id)
	return err
}

// Close statments used by [s].
func (s *SQLiteStore) Close() error {
	s.removeStmt.Close()
	s.updateStmt.Close()
	s.getStmt.Close()
	s.addStmt.Close()
	return nil
}

// NewSQLiteStore returns a SQLiteStore using the file at [filepath] and table name [table].
// If no table [table] exists, NewSQLiteStore will create one.
func NewSQLiteStore(db *sql.DB, table string) (*SQLiteStore, error) {
	s := &SQLiteStore{}
	var err error

	fields := "Title, Details, Author, Content"
	fieldsHolder := "?, ?, ?, ?"
	fieldsUpdater := "Title = ?, Details = ?, Author = ?, Content = ?"

	s.addStmt, err = db.Prepare("INSERT INTO " + table + " (ID , " + fields + ") VALUES (?, " + fieldsHolder + ");")
	if err != nil {
		return nil, err
	}

	s.getStmt, err = db.Prepare("SELECT " + fields + " FROM " + table + " WHERE ID = ? LIMIT 1;")
	if err != nil {
		return nil, err
	}

	s.updateStmt, err = db.Prepare("Update " + table + " SET " + fieldsUpdater + " WHERE ID = ?;")
	if err != nil {
		return nil, err
	}

	s.removeStmt, err = db.Prepare("DELETE FROM " + table + " WHERE ID = ?;")
	if err != nil {
		return nil, err
	}

	return s, nil
}
