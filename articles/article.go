package articles

import (
	"errors"
	"strconv"
	"strings"
	"time"
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

// Author stores information about an author.
type Author struct {
	User       string
	Name       string
	URL        string
	Permission AuthorPermission
}

// AuthorPermission is an int used to dedermine what an author is allowed to do.
type AuthorPermission int

const (
	// AuthorPermissionDisable - permission is either unknown due to an error or author is banned.
	AuthorPermissionDisable AuthorPermission = 0
	// AuthorPermissionView - Can't make suggestings, edits or create new articles.
	AuthorPermissionView AuthorPermission = 1
	//AuthorPermissionSuggest - Can make suggestings but can't make edits or create new articles.
	AuthorPermissionSuggest AuthorPermission = 2
	//AuthorPermissionLimited - Can make suggestings and edits on self made articles but can't create new articles.
	AuthorPermissionLimited AuthorPermission = 3
	//AuthorPermissionNormal - Can make suggestings, edits on self made articles and create new articles under own name.
	AuthorPermissionNormal AuthorPermission = 4
	//AuthorPermissionAdmin - Can make suggestings, edits or create new articles even once under anothers name.
	AuthorPermissionAdmin AuthorPermission = 5
)

// TimeStamp TODO: Comment on TimeStamp.
type TimeStamp struct {
	Year  int
	Month time.Month
	Day   int
}

// ErrFormat is error thrown when string is not properly formatted.
var ErrFormat = errors.New("string not in YYYY-MM-DD format")

// NewTimeStampString returns a new TimeStamp from a string in YYYY-MM-DD format.
func NewTimeStampString(s string) (TimeStamp, error) {
	sub := strings.Split(s, "-")
	if len(sub) != 3 {
		return TimeStamp{}, ErrFormat
	}

	y, err := strconv.Atoi(sub[0])
	if err != nil {
		return TimeStamp{}, ErrFormat
	}

	mInt, err := strconv.Atoi(sub[1])
	if err != nil {
		return TimeStamp{}, ErrFormat
	}
	m := time.Month(mInt)

	d, err := strconv.Atoi(sub[2])
	if err != nil {
		return TimeStamp{}, ErrFormat
	}

	return TimeStamp{
		Year:  y,
		Month: m,
		Day:   d,
	}, nil
}

// String returns a string in YYYY-MM-DD format.
func (t *TimeStamp) String() string {
	return strconv.Itoa(t.Year) + "-" + strconv.Itoa(int(t.Month)) + "-" + strconv.Itoa(t.Day)
}
