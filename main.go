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

// InstallFlag is the systemd installer flag.
var InstallFlag *bool

func init() {
	InstallFlag = flag.Bool("systemd", false, "Install systemd service 💾 ")
}

func main() {
	flag.Parse()

	// Handles installing systemd service if InstallFlag is set.
	if *InstallFlag {
		log.Println("💾  Installing systemd service:")
		path, err := os.Executable()
		if err != nil {
			log.Println("\t❌  Failed: Unable to find path: " + err.Error())
			os.Exit(1)
		}
		file := []byte("\n[Unit]\nDescription=Computing Fun web server\n[Service]\nExecStart=" + path + "\nWorkingDirectory=" + filepath.Dir(path) + "\n[Install]\nWantedBy=multi-user.target")
		err = ioutil.WriteFile("/etc/systemd/system/cf-www.service", file, 0664)
		if err != nil {
			log.Println("\t❌  Failed: Unable to write file: " + err.Error())
			os.Exit(1)
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
	router.ServeFiles("/client/*filepath", client.NewHTTPFileSystemLogErr())
	router.NotFound = http.HandlerFunc(NotFoundHandler)
	router.PanicHandler = PanicHandler

	if err := firestoredb.Init(context.TODO(), "credentials.json"); err != nil {
		log.Fatalln("🔑  " + err.Error())
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
