package html

import (
	"io"
	"net/http"
)

// ErrMessage is an Error that can be used with html.Err.
type ErrMessage struct {
	Message string
	Code    string
	Title   string
}

func (e *ErrMessage) Error() string {
	return e.Message + " : " + e.Code
}

//----- 404 Not Found -----

// NewNotFoundErrMessage is an ErrMessage for "Not Found" pages.
func NewNotFoundErrMessage() *ErrMessage {
	return &ErrMessage{
		Message: "We're sorry but we couldn't find the page you're looking for :(",
		Code:    "Error - HTTP 404: Not Found",
		Title:   "Not it",
	}
}

// NotFound - Err with NotFoundErrMessage
func NotFound(w io.Writer) (int, error) {
	return Err(NewNotFoundErrMessage(), w)
}

// NotFoundHandler responses with the NotFound error page (404 status code).
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	NotFound(w)
}

//----- 500 Internal Server Error -----

// NewPanicErrMessage is an ErrMessage for "Panic" or "Server Error" pages.
func NewPanicErrMessage() *ErrMessage {
	return &ErrMessage{
		Message: "There seems to be something wrong. Don't panic, we're already doing that for you.",
		Code:    "Error - HTTP 500: Internal Server Error",
		Title:   "Panicing",
	}
}

// Panic - Err with PanicErrMessage
func Panic(w io.Writer) (int, error) {
	return Err(NewPanicErrMessage(), w)
}

// PanicHandler responses with the Panic error page (500 status code).
func PanicHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	Panic(w)
}

//----- 503 StatusServiceUnavailable -----

// NewUnavailableErrMessage is an ErrMessage for pages that are under mantains.
func NewUnavailableErrMessage() *ErrMessage {
	return &ErrMessage{
		Message: "This page isn't ready just yet.",
		Code:    "Error - HTTP 503: Service Unavailable",
		Title:   "Work in process",
	}
}

// Unavailable - Err with UnavailableErrMessage
func Unavailable(w io.Writer) (int, error) {
	return Err(NewUnavailableErrMessage(), w)
}

// UnavailableHandler responses with the Unavailable error page (503 status code).
func UnavailableHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusServiceUnavailable)
	Unavailable(w)
}
