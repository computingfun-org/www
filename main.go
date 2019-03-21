package main

//go:generate go generate ./client

import (
	"context"
	"log"
	"net/http"

	"golang.org/x/crypto/acme/autocert"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	fsClient := NewFirestoreClientFatal(context.TODO(), "credentials.json")

	cert := autocert.Manager{
		Cache:      NewFirestoreCacheFatal(fsClient, "certs"),
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("www.computingfun.org"),
		Email:      "security@computingfun.org",
	}

	server := http.Server{
		Handler:   NewRouter(),
		TLSConfig: cert.TLSConfig(),
	}

	// DEGUG SERVER
	//log.Fatalln(server.ListenAndServe())

	go func() {
		err := http.ListenAndServe("", cert.HTTPHandler(nil))
		log.Fatalln(err)
	}()

	err := server.ListenAndServeTLS("", "")
	log.Fatalln(err)
}
