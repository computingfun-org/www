package articles

// Article is a record from an Store.
type Article struct {
	Title   string
	Details string
	Author  string
	Content string
}

// Store is a storage interface for saving and getting Article records.
type Store interface {
	Add(id string, a Article) error
	Get(id string) (Article, error)
	Update(id string, a Article) error
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
