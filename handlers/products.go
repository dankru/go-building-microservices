package handlers

import (
	"log"
	"net/http"
	"web-service-gin/go-api-gateway/data"
)

type Products struct {
	l *log.Logger
}


func NewProduct (l *log.Logger) *Products {
	return &Products{l}
}

// we need serveHTTP method for each handler to make it an implementation of handler interface
func (p *Products) ServeHTTP (rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unnable to marshall json", 500)
	}
}