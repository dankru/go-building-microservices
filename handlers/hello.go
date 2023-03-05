package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// We already know that handleFunc takes a function as a parameter and
// converts it to a http handler interface and then adds it to the server, the
// handler signature from documentation is
// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }
// so to make a handler we only need to make a struct which implements a single function
// ServeHTTP(ResponseWriter, *Request)

type Hello struct {
	// 2 so we add a field which implements a logger
	l *log.Logger
}

// 3 which we'll be able to replace later with what we want
// that is called dependency injection
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}

// That signature is what we need to satisfy Handler interface
// ResponseWriter is an interface which is used by http handler to construct http response
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// 1 We were using a simple loger here. That in itself is okay, but later we'll want some
	// control over this logging. For example, if we have a database down here, 
	// we'll want to have it as an object. We don't want to have to create
	// concrete objects in a handler if it's possible to avoid it. And the 
	// reason for this is that we want to later have more testibility and
	// be able to easily replace those things, e.g. use 'dependency injection'
	h.l.Println("Hello world")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// WriteHeader allows us to specify http code that we'll return back to the user
		// rw.WriteHeader(400)
		// Write writes the data to the connection as part of an HTTP reply.
		// rw.Write([]byte("Oops"))
		// We can perform the same tasks in one line of code with http.error
		http.Error(rw, "something went wrong", 400)
		return
	}
	// Sinice ResponseWriter interface's method Write() is also implements io.writer interface,
	// we can use it in Fprintf
	fmt.Fprintf(rw, "hello %s", d)
}

