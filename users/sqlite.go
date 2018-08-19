package admin

import (
	"database/sql"
)

// SQLiteStore is an admin.Store that uses a Sqlite database.
// Should not be made directly. Instead use NewSQLiteStore.
type SQLiteStore struct {
	hashCost    int
	addStmt     *sql.Stmt
	removeStmt  *sql.Stmt
	getPassStmt *sql.Stmt
	setPassStmt *sql.Stmt
	getNameStmt *sql.Stmt
	setNameStmt *sql.Stmt
}

// Close store and all connections/statments.
func (s *SQLiteStore) Close() error {
	s.setNameStmt.Close()
	s.getNameStmt.Close()
	s.setPassStmt.Close()
	s.getPassStmt.Close()
	s.removeStmt.Close()
	s.addStmt.Close()
	return nil
}

// Add creates a new admin user and returns error if user can't be created.
func (s *SQLiteStore) Add(user, pass string) error {
	hash, err := GeneratePassHash(user, pass, s.hashCost)
	if err == nil {
		_, err = s.addStmt.Exec(user, hash)
	}
	return err
}

// Remove deletes an admin user and returns error if unable to remove the user.
func (s *SQLiteStore) Remove(user string) error {
	_, err := s.removeStmt.Exec(user)
	return err
}

// CheckPass compares the admin user's passphrase (password) to the stored hash.
// If passphrase doesn't match the stored hash CheckPass returns an error.
// If no errors occurs and the passphrase matches the hash CheckPass returns nil.
func (s *SQLiteStore) CheckPass(user, pass string) error {
	var hash []byte
	err := s.getPassStmt.QueryRow(user).Scan(&hash)
	if err == nil {
		err = CheckPassHash(user, pass, hash)
	}
	return err
}

// ChangePass changes the admin user's passphrase (password).
func (s *SQLiteStore) ChangePass(user, pass string) error {
	hash, err := GeneratePassHash(user, pass, s.hashCost)
	if err == nil {
		_, err = s.setPassStmt.Exec(hash, user)
	}
	return err
}

// GetName returns the admin user's Name.
func (s *SQLiteStore) GetName(user string) (string, error) {
	var name string
	err := s.getNameStmt.QueryRow(user).Scan(&name)
	return name, err
}

// SetName changes the admin user's Name.
func (s *SQLiteStore) SetName(user, name string) error {
	_, err := s.setNameStmt.Exec(name, user)
	return err
}

// NewSQLiteStore returns a new SQLiteStore.
func NewSQLiteStore(db *sql.DB, table string, hashCost int) (*SQLiteStore, error) {
	s := &SQLiteStore{hashCost: hashCost}

	_, err := db.Exec("CREATE TABLE IF NOT EXISTS " + table + " (User text PRIMARY KEY NOT NULL CHECK (User != \"\"), Pass blob NOT NULL CHECK (Pass != \"\"), Name text, Permissions blob);")
	if err != nil {
		return nil, err
	}

	s.addStmt, err = db.Prepare("INSERT INTO " + table + " (User, Pass) VALUES (?, ?);")
	if err != nil {
		return nil, err
	}

	s.removeStmt, err = db.Prepare("DELETE FROM " + table + " WHERE User = ?;")
	if err != nil {
		return nil, err
	}

	s.getPassStmt, err = db.Prepare("SELECT Pass FROM " + table + " WHERE User = ? LIMIT 1;")
	if err != nil {
		return nil, err
	}

	s.setPassStmt, err = db.Prepare("Update " + table + " SET Pass = ? WHERE User = ?;")
	if err != nil {
		return nil, err
	}

	s.getNameStmt, err = db.Prepare("SELECT Name FROM " + table + " WHERE User = ? LIMIT 1;")
	if err != nil {
		return nil, err
	}

	s.setNameStmt, err = db.Prepare("Update " + table + " SET Name = ? WHERE User = ?;")
	if err != nil {
		return nil, err
	}

	return s, nil
}
