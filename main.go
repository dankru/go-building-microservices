package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"web-service-gin/go-api-gateway/handlers"

	"golang.org/x/net/context"
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

	// Now we will create our own golang server so we can tune it more precisely
	// server is a type which has some fields: Addr, Handler, TLSConfig, ReadTimeout, WriteTimeout...
	// for example, if we want to send a large file to the server, then we should have a large WriteTimeout
	// but what if we have to separate handlers, one of which needs a big timeout and another one needs small amout of time?
	// we will come to that in the next session! Because it is not that straightforward
	// IdleTimeout parameter. When we esablish connection between client and a server, there are a lot of things
	// going under the hood e.g dns lookups, handshakes etc. Those things cost time. And establishing connection for the same
	// client every time is not efficient, so we can reuse the same connection over and over again, by increasing the IdleTimeout
	// which also reduces the amout of accessible connections for clients. This is particulalry useful when you have
	// lots of microservices connected to each other. In this project we want to maintain persistent connection between those.
	// especially since we are using TLS, which is more expensive in establishing connections 

	// We are tuning our server here. Pretty self-explanatory
	s := &http.Server{
		Addr: ":3000",
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1*time.Second,
		WriteTimeout: 1*time.Second,
	}

	go func() {
		// Now we should run our server
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	// signal.Notify will broadcast a message on sigChan whenever the os recieves a kill or interrupt command
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <- sigChan
	l.Println("Recieved terminate, gratefull shutdown", sig)

	// Argument for Shutdown. It says here: I want you to wait 30 seconds to attempt to gracefully shutdown
	// if 30 seconds has passed and the requests are still not handled, shutdown forcefully
	tc, _ := context.WithTimeout(context.Background(), 30 * time.Second)
	// Basically a graceful shutdown. The server will wait untill current requests are completed and then it will shut
	s.Shutdown(tc)

	// now we should register this handler to our server
	// we bind our serveMux handler to our server	
	// http.ListenAndServe(":3000", sm)
}