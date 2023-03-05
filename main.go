package main

import (
	"log"
	"net/http"
	"os"
	"web-service-gin/go-api-gateway/handlers"
)

func main() {
	// let's make a handler
	// first we'll make a logger which we'll pass into our NewHello constructor 
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)
	// now to register this handler into our serveMux, we need to create the serveMux
	sm := http.NewServeMux()
	// and add a handler to it
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	// now we should register this handler to our server
	// we bind our serveMux handler to our server	
	http.ListenAndServe(":3000", sm)
}