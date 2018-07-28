package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gitlab.com/computingfun/computingfun.org/html"
)

func main() {
	router := httprouter.New()
	router.NotFound = newNotFoundHandler()
	router.GET("/", indexHandler)
	router.ServeFiles("/client/*filepath", http.Dir("./client"))

	/*
		cert := autocert.Manager{
			Cache:      autocert.DirCache("autocert"),
			Prompt:     autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist("computingfun.org", "www.computingfun.org"),
		}
	*/

	server := http.Server{
		Handler: router,
		//TLSConfig: &tls.Config{GetCertificate: cert.GetCertificate},
	}

	//go http.ListenAndServe("", cert.HTTPHandler(nil))
	//log.Fatalln(server.ListenAndServeTLS("", ""))
	log.Fatalln(server.ListenAndServe())
}

func newNotFoundHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		html.NotFound(w)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	html.Index(w)
}
