package main

//go:generate go generate ./client

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
	"gitlab.com/computingfun/www/client"
	"gitlab.com/computingfun/www/firestoredb"
	"gitlab.com/zacc/autocertcache"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	installFlag := flag.Bool("install", false, "Install systemd service 💾 ")
	flag.Parse()

	if *installFlag {
		log.Println("💾  Installing systemd service:")
		err := installService()
		if err != nil {
			log.Fatalln("\t❌  Failed: " + err.Error())
		}
		log.Println("\t✔️  Success")
		os.Exit(0)
	}

	router := httprouter.New()
	router.GET("/", IndexHandler)
	router.GET("/articles/", UnavailableHandler)
	router.GET("/articles/:id", ArticleHandler)
	router.GET("/games/", UnavailableHandler)
	router.GET("/games/:id", UnavailableHandler)
	if fs, err := client.NewHTTPFileSystem(); err != nil {
		router.ServeFiles("/client/*filepath", fs)
	}
	router.NotFound = http.HandlerFunc(NotFoundHandler)
	router.PanicHandler = PanicHandler

	if err := firestoredb.Init(context.TODO(), "credentials.json"); err != nil {
		log.Fatalln(err)
	}

	cert := autocert.Manager{
		Cache:      autocertcache.NewFirestore(firestoredb.AutoCertCache),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("www.computingfun.org", "beta.computingfun.org"),
		Email:      "security@computingfun.org",
	}

	server := http.Server{
		Handler:   router,
		TLSConfig: cert.TLSConfig(),
	}

	log.Println("🌐")

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
	go log.Println("⚠️  Panic: ", e, " | Request: ", r)
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

func installService() error {
	path, err := os.Executable()
	if err != nil {
		return err
	}

	file := []byte("\n[Unit]\nDescription=Computing Fun web server.\n[Service]\nExecStart=" + path + "\nWorkingDirectory=" + filepath.Dir(path) + "\n[Install]\nWantedBy=multi-user.target")
	return ioutil.WriteFile("/etc/systemd/system/cf-www.service", file, 0664)
}
