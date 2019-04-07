package client

//go:generate go get github.com/valyala/quicktemplate/qtc
//go:generate qtc

//go:generate go get github.com/rakyll/statik
//go:generate statik -src=filesystem -dest=../ -p=client

import (
	"log"
	"net/http"

	"github.com/rakyll/statik/fs"
)

// NewHTTPFileSystem TODO: comment client.NewHTTPFileSystem
func NewHTTPFileSystem() (http.FileSystem, error) {
	return fs.New()
}

// NewHTTPFileSystemLogErr TODO: comment client.NewHTTPFileSystemLogErr
func NewHTTPFileSystemLogErr() http.FileSystem {
	fs, err := NewHTTPFileSystem()
	if err != nil {
		log.Println("⚠️  " + err.Error())
	}
	return fs
}
