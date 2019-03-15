package main

import (
	"database/sql"
	"log"
	"net/http"

	"gitlab.com/zacc/autocertcache"
	"golang.org/x/crypto/acme/autocert"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./cf.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	/*
		ArticleStore, err = articles.NewSQLiteStore(db, "Articles")
		if err != nil {
			log.Fatalln(err)
		}
		defer ArticleStore.Close()
	*/

	certCache, err := autocertcache.NewSQLite(db, "Certs")
	if err != nil {
		log.Fatalln(err)
	}
	defer certCache.Close()

	cert := autocert.Manager{
		Cache:      certCache,
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("www.computingfun.org"),
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

	err = server.ListenAndServeTLS("", "")
	log.Fatalln(err)
}
