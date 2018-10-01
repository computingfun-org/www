package html

import (
	"bytes"
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

// NewNotFoundErrMessage is an ErrMessage for "Not Found" pages.
func NewNotFoundErrMessage() *ErrMessage {
	return &ErrMessage{
		Message: "We're sorry but we couldn't find the page you're looking for :(",
		Code:    "Error - HTTP 404: Not Found",
		Title:   "Not it (404)",
	}
}

// NotFound - Err with NotFoundErrMessage
func NotFound(buffer *bytes.Buffer) {
	Err(NewNotFoundErrMessage(), buffer)
}

// NewPanicErrMessage is an ErrMessage for "Panic" or "Server Error" pages.
func NewPanicErrMessage() *ErrMessage {
	return &ErrMessage{
		Message: "There seems to be something wrong. Don't panic, we're already doing that for you.",
		Code:    "Error - HTTP 500: Internal Server Error",
		Title:   "Panicing (500)",
	}
}

// Panic - Err with PanicErrMessage
func Panic(buffer *bytes.Buffer) {
	Err(NewPanicErrMessage(), buffer)
}

// NewComingSoonErrMessage is an ErrMessage for pages that aren't yet available.
func NewComingSoonErrMessage() *ErrMessage {
	return &ErrMessage{
		Message: "We aren't ready just yet. In the meantime you can check out the links to the left :P",
		Code:    "Coming soon, to the interweb.",
		Title:   "Coming Soon!",
	}
}

// ComingSoon - Err with ComingSoonErrMessage
func ComingSoon(buffer *bytes.Buffer) {
	Err(NewComingSoonErrMessage(), buffer)
}
