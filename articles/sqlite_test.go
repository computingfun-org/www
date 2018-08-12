package articles

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

// TestSQLiteStore use to test common use cases for SQLiteStore.
func TestSQLiteStore(t *testing.T) {
	defer os.Remove("./test.db")
	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	se, err := NewSQLiteStore(db, "")
	if err == nil {
		t.Fatal("Table name shouldn't allow empty sting.")
		se.Close()
	}

	s, err := NewSQLiteStore(db, "TestArticles")
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()

	a1 := &Article{Title: "Title1", Details: "Details1", Content: "Content1"}
	a2 := &Article{Title: "Title2", Details: "Details2", Content: "Content2"}
	e1 := &Article{Title: "ERROR", Details: "ERROR", Content: "ERROR"}

	err = s.Add(a1, "ID1")
	if err != nil {
		t.Error(err)
	}

	err = s.Add(a2, "ID2")
	if err != nil {
		t.Error(err)
	}

	err = s.Add(e1, "ID1")
	if err == nil {
		t.Error("Added ID1 for a secound time supposed to error")
	}

	getA1, err := s.Get("ID1")
	if err != nil {
		t.Error(err)
	}

	if getA1.Title != a1.Title {
		t.Error("Get not returning proper 'Title'")
	}

	if getA1.Details != a1.Details {
		t.Error("Get not returning proper 'Details'")
	}

	if getA1.Content != a1.Content {
		t.Error("Get not returning proper 'Contents'")
	}

	err = s.Remove("ID1")
	if err != nil {
		t.Error(err)
	}

	e2, err := s.Get("ID1")
	if err == nil || e2 != nil {
		t.Error("ID1 being receive after being removed supposed to error")
	}

	e2, err = s.Get("ID3")
	if err == nil || e2 != nil {
		t.Error("ID3 being receive supposed to error")
	}

	err = s.Add(&Article{Title: "T", Details: "D", Content: "C"}, "")
	if err == nil {
		t.Error("Empty sting shouldn't be allowed as id")
	}

	err = s.Close()
	if err != nil {
		t.Error(err)
	}
}
