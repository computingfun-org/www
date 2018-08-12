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

// ArticleStore is the storage for articles.
var ArticleStore articles.Store

func main() {
	handler := httprouter.New()
	handler.PanicHandler = PanicHandler
	handler.NotFound = GetNotFoundHandler()
	handler.GET("/", IndexHandler)
	handler.GET("/articles/:id", ArticleHandler)
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

// NotFoundHandler responses with the NotFound error page (404 status code).
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	html.NotFound(w)
}

// GetNotFoundHandler just returns NotFoundHandler.
// Used because httprouter rounter's NotFound won't take NotFoundHandler without this getter func.
func GetNotFoundHandler() http.HandlerFunc {
	return NotFoundHandler
}

// PanicHandler responses with the Panic error page (500 status code) and logs its.
func PanicHandler(w http.ResponseWriter, r *http.Request, i interface{}) {
	log.Println("Panic: ", i, " | Request: ", r, " | Response: ", w)
	w.WriteHeader(http.StatusInternalServerError)
	html.Panic(w)
}

// IndexHandler responses with the home page.
func IndexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	html.ComingSoon(w)
}

// ArticleHandler responses with an article page for the article with id [:id].
// If article is not found ArticleHandler responses with NotFoundHandler.
func ArticleHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	a, err := ArticleStore.Get(p.ByName("id"))
	if err != nil {
		NotFoundHandler(w, r)
		return
	}
	html.Article(a, w)
}
