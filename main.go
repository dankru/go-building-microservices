package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	// 3 HandleFunc registers a function to a path on a thing called the default serve mux
	// 4 serve mux is an http handler
	// 5 HandleFunc is a convinience method. It takes my function and makes a handler from it, and
	// 6 then adds this handler to the default serve mux
	http.HandleFunc("/", getIndexPage)
	http.HandleFunc("/goodbye", goodbye)

	// 1 ListenAndServe constructs an http server and registers a default handler to it
	// 2 nil handler means that ListenAndServe will use a default serve mux
	// 7 default serve mux has handlers that we've already added to it
	// 8 it determines which function will be executed when we request our path
	http.ListenAndServe(":3000", nil)

}

//ResponseWriter is an interface which is used by http handler to construct http response
func getIndexPage(rw http.ResponseWriter, r *http.Request) {
	log.Println("hello world")
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
func goodbye(rw http.ResponseWriter, r *http.Request) {
	log.Println("goodbye")
}