package client

//go:generate go get github.com/rakyll/statik
//go:generate statik -src=./ -dest=../ -p=client

import (
	"net/http"

	"github.com/rakyll/statik/fs"
)

// NewHTTPFileSystem TODO: comment client.NewHTTPFileSystem
func NewHTTPFileSystem() (http.FileSystem, error) {
	return fs.New()
}
