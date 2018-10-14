package articles

import (
	"errors"
	"strconv"
	"strings"
)

// Article is a record from an Store.
type Article struct {
	Title     string
	Details   string
	Author    string
	TimeStamp TimeStamp
	Content   string
}

// Store is a storage interface for saving and getting Article records.
type Store interface {
	Add(id string, a *Article) error
	Get(id string) (*Article, error)
	Update(id string, a *Article) error
	Remove(id string) error
	Close() error
}

// TimeStamp TODO: Comment on TimeStamp.
type TimeStamp struct {
	Year  int
	Month int
	Day   int
}

// ErrTimeFormat is error thrown when string is not properly formatted.
var ErrTimeFormat = errors.New("string not in [Year]-[Month]-[Day] format")

// NewTimeStamp returns a new TimeStamp from a string in [Year]-[Month]-[Day] format.
func NewTimeStamp(s string) TimeStamp {
	ts := TimeStamp{}
	sub := strings.Split(s, "-")
	l := len(sub)

	if l > 0 {
		ts.Year, _ = strconv.Atoi(sub[0])
	}
	if l > 1 {
		ts.Month, _ = strconv.Atoi(sub[1])
	}
	if l > 2 {
		ts.Day, _ = strconv.Atoi(sub[2])
	}

	return ts
}

// String returns a string of TimeStamp in [Year]-[Month]-[Day] format.
func (t *TimeStamp) String() string {
	return strconv.Itoa(t.Year) + "-" + strconv.Itoa(int(t.Month)) + "-" + strconv.Itoa(t.Day)
}
