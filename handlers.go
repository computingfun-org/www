package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gitlab.com/computingfun/www/client"
)

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
	go log.Println("🛑  Panic: ", e, " | Request: ", r)
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
	panic("ArticleHandler is not ready yet.")
}
