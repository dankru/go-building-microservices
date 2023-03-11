package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
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
	// handling requests
	if r.Method == http.MethodGet {
		p.GetProducts(rw, r)
		return
	}		
	
	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		// implementing picking up id from the URI
		// expect the id in the URI
		// r is a really simple regex
		reg := regexp.MustCompile(`/([0-9+])`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			p.l.Println("Invalid URI more than one id")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("Invalid URI more than one capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "invalid uri failed to convert to number", 500)
		}

		p.l.Println("got id", id)
	}
	
	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET products")

	// fetch the products from the data store
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unnable to marshall json", 500)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST products")

	prod := &data.Product{}
	
	err := prod.FromJSON(r.Body)
	if err != nil {
			http.Error(rw, "Unable to unmarshall JSON", http.StatusBadRequest)
	}	

	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}