package admin

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Store TODO: Comment
type Store interface {
	Close() error

	Add(user, pass string) error
	Remove(user string) error

	CheckPass(user, pass string) error
	ChangePass(user, pass string) error

	GetName(user string) (string, error)
	SetName(user, name string) error
}

func salter(user, pass string) []byte {
	return []byte(pass + user)
}

var (
	// PassRequiredLen is the lenght a passphrase (password) must be to be vaild.
	PassRequiredLen = 10

	// ErrEmptyUser is the error that is returned when trying to use an empty string as a user.
	ErrEmptyUser = errors.New("user can not be an empty string")

	// ErrShortPass is the error that is returned when trying to use a passphrase (password) that is too short.
	ErrShortPass = errors.New("pass is too short")
)

// GeneratePassHash returns a hash.
func GeneratePassHash(user, pass string, cost int) ([]byte, error) {
	if user == "" {
		return nil, ErrEmptyUser
	}
	if len(pass) < PassRequiredLen {
		return nil, ErrShortPass
	}
	return bcrypt.GenerateFromPassword(salter(user, pass), cost)
}

// CheckPassHash checks if user and pass matches the hash.
func CheckPassHash(user, pass string, hash []byte) error {
	return bcrypt.CompareHashAndPassword(hash, salter(user, pass))
}
