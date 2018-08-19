package admin

import (
	"database/sql"
	"encoding/json"
	"errors"

	"gitlab.com/computingfun/computingfun.org/admin/permission"

	"golang.org/x/crypto/bcrypt"
)

// ErrEmptyPass is the error that is returned when trying to use an empty string as a passphrase (password).
var ErrEmptyPass = errors.New("pass can not be an empty string")

// ErrPermissionDenied is the error that is returned when checking for PermissionLevel of a Permission and
// it fails because the admin user's PermissionLevel isn't high enough.
var ErrPermissionDenied = errors.New("admin user doesn't have permission")

// SQLiteStore is an admin.Store that uses a Sqlite database.
// Should not be made directly. Instead use NewSQLiteStore.
type SQLiteStore struct {
	hashCost           int
	addStmt            *sql.Stmt
	removeStmt         *sql.Stmt
	getPassStmt        *sql.Stmt
	setPassStmt        *sql.Stmt
	getNameStmt        *sql.Stmt
	setNameStmt        *sql.Stmt
	getPermissionsStmt *sql.Stmt
	setPermissionsStmt *sql.Stmt
}

// Close store and all connections/statments.
func (s *SQLiteStore) Close() error {
	s.setPermissionsStmt.Close()
	s.getPermissionsStmt.Close()
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
	hash, err := bcrypt.GenerateFromPassword([]byte(pass+user), s.hashCost)
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
		err = bcrypt.CompareHashAndPassword(hash, []byte(pass+user))
	}
	return err
}

// ChangePass changes the admin user's passphrase (password).
func (s *SQLiteStore) ChangePass(user, pass string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass+user), s.hashCost)
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

// GetPermissionMap returns a permission.Map of admin user's permissions.
func (s *SQLiteStore) GetPermissionMap(user string) (permission.Map, error) {
	var p permission.Map
	var raw []byte
	err := s.getPermissionsStmt.QueryRow(user).Scan(&raw)
	if err == nil && len(raw) > 0 {
		err = json.Unmarshal(raw, &p)
	}
	if p == nil {
		p = permission.NewMap()
	}
	return p, err
}

// SetPermissionMap changes admin user's permissions to match a permission.Map.
func (s *SQLiteStore) SetPermissionMap(user string, permissions permission.Map) error {
	raw, err := json.Marshal(permissions)
	if err == nil {
		_, err = s.setPermissionsStmt.Exec(raw, user)
	}
	return err
}

// GetPermission returns the PermissionLevel of [permission] the admin user has.
func (s *SQLiteStore) GetPermission(user string, permission permission.Type) (permission.Level, error) {
	p, err := s.GetPermissionMap(user)
	return p[permission], err
}

// SetPermission changes the PermissionLevel of [permission] for the admin user.
func (s *SQLiteStore) SetPermission(user string, permission permission.Type, level permission.Level) error {
	p, err := s.GetPermissionMap(user)
	if err == nil {
		p[permission] = level
		err = s.SetPermissionMap(user, p)
	}
	return err
}

// CheckPermission returns nil if admin user has at least [level] of [permission].
func (s *SQLiteStore) CheckPermission(user string, permission permission.Type, level permission.Level) error {
	l, err := s.GetPermission(user, permission)
	if err == nil && l < level {
		return ErrPermissionDenied
	}
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

	s.getPermissionsStmt, err = db.Prepare("SELECT Permissions FROM " + table + " WHERE User = ? LIMIT 1;")
	if err != nil {
		return nil, err
	}

	s.setPermissionsStmt, err = db.Prepare("Update " + table + " SET Permissions = ? WHERE User = ?;")
	if err != nil {
		return nil, err
	}

	return s, nil
}
