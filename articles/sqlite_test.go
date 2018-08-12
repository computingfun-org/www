package articles

import (
	"os"
	"testing"
)

// TestSQLiteStore use to test common use cases for SQLiteStore.
func TestSQLiteStore(t *testing.T) {
	TestDBPath := "./test.db"
	os.Remove(TestDBPath)

	s, err := NewSQLiteStore("TestArticles", TestDBPath)
	if err != nil {
		t.Fatal(err)
	}
	defer s.Close()
	defer os.Remove(TestDBPath)

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

	err = s.Close()
	if err != nil {
		t.Error(err)
	}

	err = os.Remove(TestDBPath)
	if err != nil {
		t.Error(err)
	}
}
