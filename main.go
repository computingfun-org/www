package main

import (
	"crypto/tls"
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gitlab.com/computingfun/www/articles"
	"gitlab.com/computingfun/www/html"
	"golang.org/x/crypto/acme/autocert"

	_ "github.com/mattn/go-sqlite3"
)

var (
	// ArticleStore is the storage for articles.
	ArticleStore *articles.SQLiteStore
)

func main() {
	// Database
	db, err := sql.Open("sqlite3", "./cf.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	// Article table
	ArticleStore, err = articles.NewSQLiteStore(db, "Articles")
	if err != nil {
		log.Fatalln(err)
	}
	defer ArticleStore.Close()

	// Router
	handler := httprouter.New()
	handler.GET("/", UnavailableHandler)
	handler.GET("/articles/", UnavailableHandler)
	handler.GET("/articles/:id", ArticleHandler)
	handler.GET("/games/", UnavailableHandler)
	handler.GET("/games/:id", UnavailableHandler)
	handler.ServeFiles("/client/*filepath", http.Dir("./client"))
	handler.NotFound = http.HandlerFunc(html.NotFoundHandler)
	handler.PanicHandler = func(w http.ResponseWriter, r *http.Request, e interface{}) {
		go log.Println("Panic: ", e, " | Request: ", r)
		html.PanicHandler(w, r)
	}

	// TLS certificate
	cert := autocert.Manager{
		Cache:      autocert.DirCache("autocert"),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("computingfun.org", "www.computingfun.org"),
	}

	// Server
	server := http.Server{
		Handler:   handler,
		TLSConfig: &tls.Config{GetCertificate: cert.GetCertificate},
	}

	// HTTP server, handles Let's Encrypt challenge responses and HTTP redirects.
	go func() {
		err := http.ListenAndServe("", cert.HTTPHandler(nil))
		log.Fatalln(err)
	}()

	err = server.ListenAndServeTLS("", "")
	log.Fatalln(err)
}

// UnavailableHandler is an adapter for html.UnavailableHandler.
func UnavailableHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	html.UnavailableHandler(w, r)
}

// ArticleHandler responses with an article page for the article with id [:id].
// If article is not found ArticleHandler responses with NotFoundHandler.
func ArticleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	a, err := ArticleStore.Get(p.ByName("id"))
	if err != nil {
		html.NotFoundHandler(w, r)
		return
	}
	html.Article(a, w)
}
