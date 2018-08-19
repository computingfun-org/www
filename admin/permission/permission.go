package permission

// Map is a map where the key is a string representing a permission type and
// the value is an uint representing the permission level.
// A vaule of 0 should be seen as an unknown permission level.
type Map map[Type]Level

// NewMap makes and returns an empty non-nil permissions map.
func NewMap() Map {
	return make(Map)
}

// Type is a sring representing a type of permission.
type Type string

// Level is a uint representing a level of permission.
type Level uint

const (
	// Active - is admin allowed to sign in.
	Active Type = "Active"

	// ActiveNotAllow - NOT allowed to sign in.
	ActiveNotAllow Level = 0

	// ActiveAllow - allowed to sign in.
	ActiveAllow Level = 1

	// CreateArticle - can user create articles.
	CreateArticle Type = "CreateArticle"

	// CreateArticleNone - can't create articles.
	CreateArticleNone Level = 0

	// CreateArticleAllow - can create articles in own name.
	CreateArticleAllow Level = 1

	// CreateArticleAdmin - can create articles in anyones name.
	CreateArticleAdmin Level = 2

	// EditArticle - at what level can admin edit articles.
	EditArticle Type = "EditArticle"

	// EditArticleNone - can't edit or suggest edits.
	EditArticleNone Level = 0

	// EditArticleSuggest - can't edit but can suggest edits.
	EditArticleSuggest Level = 1

	// EditArticleNormal - can edit self created artices and suggest edits.
	EditArticleNormal Level = 2

	// EditArticleAdmin - can edit and suggest edits on any artices.
	EditArticleAdmin Level = 3
)
