package articles

import (
	"database/sql"
)

// SQLiteArticleStore is a SQLite store for articles.
type SQLiteArticleStore struct {
	addStmt    *sql.Stmt
	getStmt    *sql.Stmt
	updateStmt *sql.Stmt
	removeStmt *sql.Stmt
}

var _ ArticleStore = &SQLiteArticleStore{}

// Add an article [a] to the SQLite database with the ID [id].
func (s *SQLiteArticleStore) Add(a Article) error {
	_, err := s.addStmt.Exec(
		a.Author,
		a.Title,
		a.Details,
		a.Content,
		a.Created,
		a.Modified)
	return err
}

// Get an article with the id [id] from the SQLite database.
func (s *SQLiteArticleStore) Get(id string) (Article, error) {
	a := Article{}
	err := s.getStmt.QueryRow(id).Scan(
		&a.Author,
		&a.Title,
		&a.Details,
		&a.Content,
		&a.Created,
		&a.Modified)
	return a, err
}

// Update updates the article with id [id] with the data from [a].
func (s *SQLiteArticleStore) Update(id string, a Article) error {
	_, err := s.updateStmt.Exec(
		a.Author,
		a.Title,
		a.Details,
		a.Content,
		a.Created,
		a.Modified)
	return err
}

// Remove an article with the id [id] from the SQLite database.
func (s *SQLiteArticleStore) Remove(id string) error {
	_, err := s.removeStmt.Exec(id)
	return err
}

// Close statments used by [s].
func (s *SQLiteArticleStore) Close() error {
	s.removeStmt.Close()
	s.updateStmt.Close()
	s.getStmt.Close()
	s.addStmt.Close()
	return nil
}

// NewSQLiteArticleStore returns a SQLiteStore using the file at [filepath] and table name [table].
// If no table [table] exists, NewSQLiteStore will create one.
func NewSQLiteArticleStore(db *sql.DB, table string) (*SQLiteArticleStore, error) {
	s := &SQLiteArticleStore{}
	var err error

	fields := []string{"Author", "Title", "Details", "Content", "Created", "Modified"}
	fieldsLen := len(fields)

	fieldsList := ""
	fieldsHolder := ""
	fieldsUpdater := ""
	for i := 0; i < fieldsLen; i++ {
		fieldsList += fields[i]
		fieldsHolder += "?"
		fieldsUpdater += fields[i] + " = ?"
		if i != fieldsLen-1 {
			fieldsList += ", "
			fieldsHolder += ", "
			fieldsUpdater += ", "
		}
	}

	s.addStmt, err = db.Prepare("INSERT INTO " + table + " (ID , " + fieldsList + ") VALUES (?, " + fieldsHolder + ");")
	if err != nil {
		return nil, err
	}

	s.getStmt, err = db.Prepare("SELECT " + fieldsList + " FROM " + table + " WHERE ID = ? LIMIT 1;")
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
