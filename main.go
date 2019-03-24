package main

//go:generate go generate ./client

import (
	"context"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/julienschmidt/httprouter"
	"gitlab.com/computingfun/www/client"
	"gitlab.com/zacc/autocertcache"
	"golang.org/x/crypto/acme/autocert"
	"google.golang.org/api/option"
)

var (
	// AutoCertCacheCollection ...
	AutoCertCacheCollection *firestore.CollectionRef

	// UserCollection ...
	UserCollection *firestore.CollectionRef

	// AdminCollection ...
	AdminCollection *firestore.CollectionRef

	// AuthorCollection ...
	AuthorCollection *firestore.CollectionRef

	// ArticleCollection ...
	ArticleCollection *firestore.CollectionRef
)

// GetCollectionFatal ...
func GetCollectionFatal(client *firestore.Client, path string) *firestore.CollectionRef {
	c := client.Collection(path)
	if c == nil {
		log.Fatalln("GetCollectionFatal: Collection for " + path + "returned nil.")
	}
	return c
}

func main() {
	router := httprouter.New()
	router.GET("/", IndexHandler)
	router.GET("/articles/", UnavailableHandler)
	router.GET("/articles/:id", ArticleHandler)
	router.GET("/games/", UnavailableHandler)
	router.GET("/games/:id", UnavailableHandler)
	router.ServeFiles("/client/*filepath", client.NewHTTPFileSystemFatal())
	router.NotFound = http.HandlerFunc(NotFoundHandler)
	router.PanicHandler = PanicHandler

	//log.Fatalln(http.ListenAndServe("", router))

	{
		ctx := context.TODO()
		options := option.WithCredentialsFile("credentials.json")
		app, err := firebase.NewApp(ctx, nil, options)
		if err != nil {
			log.Fatalln(err)
		}
		fsclient, err := app.Firestore(ctx)
		if err != nil {
			log.Fatalln(err)
		}

		AutoCertCacheCollection = GetCollectionFatal(fsclient, "certs")
		UserCollection = GetCollectionFatal(fsclient, "users")
		AdminCollection = GetCollectionFatal(fsclient, "admins")
		AuthorCollection = GetCollectionFatal(fsclient, "authors")
		ArticleCollection = GetCollectionFatal(fsclient, "articles")
	}

	cert := autocert.Manager{
		Cache:      autocertcache.NewFirestore(AutoCertCacheCollection),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("www.computingfun.org", "beta.computingfun.org"),
		Email:      "security@computingfun.org",
	}

	server := http.Server{
		Handler:   router,
		TLSConfig: cert.TLSConfig(),
	}

	go func() {
		err := http.ListenAndServe("", cert.HTTPHandler(nil))
		log.Fatalln(err)
	}()

	err := server.ListenAndServeTLS("", "")
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
	/*
		a, err := ArticleStore.Get(p.ByName("id"))
		if err != nil {
			NotFoundHandler(w, r)
			return
		}
		client.WriteHTML(w, client.ArticlePage(a))
	*/
}
