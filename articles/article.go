package articles

// Article is a record from an Store.
type Article struct {
	Title   string
	Details string
	Content string
}

// Store is a storage interface for saving and getting Article records.
type Store interface {
	Add(a *Article, id string) error
	Get(id string) (*Article, error)
	Remove(id string) error
	Close() error
}
