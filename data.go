package main

// User ...
type User struct {
	ID          string
	Pass        string
	Email       string
	Phone       string
	DisplayName string
	FirstName   string
	LastName    string
}

// Admin ...
type Admin struct {
	ID    string // linked to User.ID
	Level uint
}

// Author of an article.
type Author struct {
	ID      string // linked to User.ID
	SiteURL string
}

// Article ...
type Article struct {
	ID       string
	AuthorID string
	Title    string
	Details  string
	Content  string
	Created  uint
	Modified uint
}
