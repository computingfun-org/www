package articles

import (
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
	User string
	Name string
	URL  string
}

type TimeStamp struct {
	Year  int
	Month time.Month
	Day   int
}

// NewTimeStampFromStringSep returns a new TimeStamp from a the string [s].
// [s] being in YYYYSMMSDD with S being [sep].
// Only returns first error to occur.
func NewTimeStampFromStringSep(s string, sep string) (TimeStamp, error) {
	t := TimeStamp{}
	var errReturn error
	sub := strings.Split(s, sep)

	t.Year, errReturn = strconv.Atoi(sub[0])
	m, err := strconv.Atoi(sub[1])
	if errReturn == nil {
		errReturn = err
	}
	t.Month = time.Month(m)
	t.Day, err = strconv.Atoi(sub[2])
	if errReturn == nil {
		errReturn = err
	}

	return t, errReturn
}

// NewTimeStampFromString returns a new TimeStamp from a the string [s].
// [s] being in YYYY-MM-DD.
func NewTimeStampFromString(s string) (TimeStamp, error) {
	return NewTimeStampFromStringSep(s, "-")
}

// YearText returns the string of Year.
func (t *TimeStamp) YearText() string {
	return strconv.Itoa(t.Year)
}

// MonthText returns the string of Month.
func (t *TimeStamp) MonthText() string {
	return strconv.Itoa(int(t.Month))
}

// DayText returns the string of Day.
func (t *TimeStamp) DayText() string {
	return strconv.Itoa(t.Day)
}

// Text returns a string in YYYYSMMSDD format with S being [s].
func (t *TimeStamp) Text(s string) string {
	return t.YearText() + s + t.MonthText() + s + t.DayText()
}

// String returns a string in YYYY-MM-DD format.
// Same as Text with "-" passed in as [sep].
func (t *TimeStamp) String() string {
	return t.Text("-")
}
