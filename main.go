package main

//go:generate go generate ./html ./client

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gitlab.com/computingfun/www/articles"
	"gitlab.com/computingfun/www/client"
	"gitlab.com/computingfun/www/html"
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
	handler.NotFound = http.HandlerFunc(html.NotFoundHandler)
	handler.PanicHandler = func(w http.ResponseWriter, r *http.Request, e interface{}) {
		go log.Println("Panic: ", e, " | Request: ", r)
		html.PanicHandler(w, r)
	}

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

// UnavailableHandler is an adapter for html.UnavailableHandler.
func UnavailableHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	html.UnavailableHandler(w, r)
}

// IndexHandler responses with the home page.
func IndexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	html.Index(w)
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
