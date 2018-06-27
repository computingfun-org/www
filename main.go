package main

import (
	"crypto/tls"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	router := httprouter.New()
	router.NotFound = notFoundHandler()
	router.GET("/", indexHandler)
	router.GET("/game", gameHandler)
	router.ServeFiles("/client/*filepath", http.Dir("./static"))

	cert := autocert.Manager{
		Cache:      autocert.DirCache("autocert"),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("computingfun.org", "www.computingfun.org"),
	}

	server := http.Server{
		Handler:   router,
		TLSConfig: &tls.Config{GetCertificate: cert.GetCertificate},
	}

	go http.ListenAndServe("", cert.HTTPHandler(nil))
	log.Fatalln(server.ListenAndServeTLS("", ""))
}

func notFoundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func indexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}

func gameHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
