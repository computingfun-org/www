package main

//go:generate go generate ./client

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gitlab.com/computingfun/www/articles"
	"gitlab.com/computingfun/www/client"
	"gitlab.com/zacc/autocertcache"
	"golang.org/x/crypto/acme/autocert"

	_ "github.com/mattn/go-sqlite3"
)

var (
	// ArticleStore is the storage for articles.
	ArticleStore *articles.SQLiteStore
)

func main() {
	db, err := sql.Open("sqlite3", "./cf.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	ArticleStore, err = articles.NewSQLiteStore(db, "Articles")
	if err != nil {
		log.Fatalln(err)
	}
	defer ArticleStore.Close()

	hfs, err := client.NewHTTPFileSystem()
	if err != nil {
		log.Fatal(err)
	}

	handler := httprouter.New()
	handler.GET("/", IndexHandler)
	handler.GET("/articles/", UnavailableHandler)
	handler.GET("/articles/:id", ArticleHandler)
	handler.GET("/games/", UnavailableHandler)
	handler.GET("/games/:id", UnavailableHandler)
	handler.ServeFiles("/client/*filepath", hfs)
	handler.NotFound = http.HandlerFunc(NotFoundHandler)
	handler.PanicHandler = PanicHandler

	certCache, err := autocertcache.NewSQL(db, "Certs")
	if err != nil {
		log.Fatalln(err)
	}
	defer certCache.Close()

	cert := autocert.Manager{
		Cache:      certCache,
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("computingfun.org", "www.computingfun.org", "beta.computingfun.org"),
	}

	server := http.Server{
		Handler:   handler,
		TLSConfig: cert.TLSConfig(),
	}

	go func() {
		err := http.ListenAndServe("", cert.HTTPHandler(nil))
		log.Fatalln(err)
	}()

	err = server.ListenAndServeTLS("", "")
	log.Fatalln(err)
}

// UnavailableHandler is an adapter for ...
func UnavailableHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusServiceUnavailable)
	client.WriteHTML(w, client.UnavailablePage)
}

// NotFoundHandler is an adapter for ...
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	client.WriteHTML(w, client.NotFoundPage)
}

// PanicHandler is an adapter for ...
func PanicHandler(w http.ResponseWriter, r *http.Request, e interface{}) {
	go log.Println("Panic: ", e, " | Request: ", r)
	w.WriteHeader(http.StatusInternalServerError)
	client.WriteHTML(w, client.PanicPage)
}

// IndexHandler responses with the home page.
func IndexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	client.WriteHTML(w, client.IndexPage{})
}

// ArticleHandler responses with an article page for the article with id [:id].
// If article is not found ArticleHandler responses with NotFoundHandler.
func ArticleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	a, err := ArticleStore.Get(p.ByName("id"))
	if err != nil {
		NotFoundHandler(w, r)
		return
	}
	client.WriteHTML(w, client.ArticlePage(*a))
}
