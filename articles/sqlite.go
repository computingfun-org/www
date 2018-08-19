package articles

import (
	"database/sql"
	"errors"
)

// SQLiteStore is a SQLite ArticleStore.
type SQLiteStore struct {
	addStmt    *sql.Stmt
	getStmt    *sql.Stmt
	removeStmt *sql.Stmt
}

// Add an article [a] to the SQLite database with the ID [id].
func (s *SQLiteStore) Add(a *Article, id string) error {
	if id == "" {
		return errors.New("SQLiteStore.Add: Can't add article with id \"\"")
	}
	_, err := s.addStmt.Exec(id, a.Title, a.Details, a.Author, a.Date, a.Content)
	return err
}

// Get an article with the id [id] from the SQLite database.
func (s *SQLiteStore) Get(id string) (*Article, error) {
	if id == "" {
		return nil, errors.New("SQLiteStore.Get: Can't get article with id \"\"")
	}
	a := &Article{}
	r := s.getStmt.QueryRow(id)
	if err := r.Scan(&a.Title, &a.Details, &a.Author, &a.Date, &a.Content); err != nil {
		return nil, err
	}
	return a, nil
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
	s.getStmt.Close()
	s.addStmt.Close()
	return nil
}

// NewSQLiteStore returns a SQLiteStore using the file at [filepath] and table name [table].
// If no table [table] exists, NewSQLiteStore will create one.
func NewSQLiteStore(db *sql.DB, table string) (*SQLiteStore, error) {
	s := &SQLiteStore{}
	var err error

	fields := "Title, Details, Author, Date, Content"
	fieldsHolder := "?, ?, ?, ?, ?"
	fieldsCreater := "Title text, Details text, Author text, Date text, Content text"
	id := "ID"
	idHolder := "?"
	idCreater := id + " text PRIMARY KEY NOT NULL CHECK (" + id + " != \"\")"

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS " + table + " (" + idCreater + ", " + fieldsCreater + ");")
	if err != nil {
		return nil, err
	}

	s.addStmt, err = db.Prepare("INSERT INTO " + table + " (" + id + ", " + fields + ") VALUES (" + idHolder + ", " + fieldsHolder + ");")
	if err != nil {
		return nil, err
	}

	s.getStmt, err = db.Prepare("SELECT " + fields + " FROM " + table + " WHERE " + id + " = " + idHolder + " LIMIT 1;")
	if err != nil {
		return nil, err
	}

	s.removeStmt, err = db.Prepare("DELETE FROM " + table + " WHERE " + id + " = " + idHolder + ";")
	if err != nil {
		return nil, err
	}

	return s, nil
}
