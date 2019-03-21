package articles

// Article is a record from an Store.
type Article struct {
	Author   string
	Title    string
	Details  string
	Content  string
	Created  uint
	Modified uint
}

// Author of an article.
type Author struct {
	User      string
	FirstName string
	LastName  string
	SiteURL   string
	PicURL    string
}

// ArticleStore is a storage interface for saving and getting Article records.
type ArticleStore interface {
	Add(a Article) error
	Get(id string) (Article, error)
	Update(id string, a Article) error
	Remove(id string) error
	Close() error
}

// AuthorStore is a storage interface for saving and getting Author records.
type AuthorStore interface {
	Add(a Author) error
	Get(user string) (Author, error)
	Update(a Author) error
	Remove(user string) error
	Close() error
}
