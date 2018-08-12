package main

import (
	"crypto/tls"
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gitlab.com/computingfun/computingfun.org/articles"
	"gitlab.com/computingfun/computingfun.org/html"
	"golang.org/x/crypto/acme/autocert"

	_ "github.com/mattn/go-sqlite3"
)

// ArticleStore is the storage for articles
var ArticleStore articles.Store

func main() {
	var err error

	handler := httprouter.New()
	handler.NotFound = NotFoundHandlerWrapper()
	handler.GET("/", IndexHandler)
	handler.GET("/article/:id", ArticleHandler)
	handler.ServeFiles("/client/*filepath", http.Dir("./client"))

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

	cert := autocert.Manager{
		Cache:      autocert.DirCache("autocert"),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("computingfun.org", "www.computingfun.org"),
	}

	server := http.Server{
		Handler:   handler,
		TLSConfig: &tls.Config{GetCertificate: cert.GetCertificate},
	}

	go http.ListenAndServe("", cert.HTTPHandler(nil))
	log.Fatalln(server.ListenAndServeTLS("", ""))
}

// NotFoundHandlerWrapper just returns NotFoundHandler.
// Used because httprouter rounter's NotFound won't take NotFoundHandler without this wrapper func.
func NotFoundHandlerWrapper() http.HandlerFunc {
	return NotFoundHandler
}

// NotFoundHandler responses with the 404 error page.
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	html.NotFound(w)
}

// IndexHandler responses with the home page.
func IndexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	html.Index(w)
}

// ArticleHandler responses with an article with the id [:id].
func ArticleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	a, err := ArticleStore.Get(p.ByName("id"))
	if err != nil {
		NotFoundHandler(w, r)
		return
	}
	log.Println(a.Title)
	html.Index(w)
}
