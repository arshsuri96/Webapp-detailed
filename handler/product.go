package handler

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/arshsuri96/test/product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.GetProducts(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.addProducts(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		p.l.Println("PUT", r.URL.Path)
		res := regexp.MustCompile(`/([0-9]+)`)
		g := res.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			p.l.Println("Invalid URL")
			http.Error(rw, "invalid URI", http.StatusBadRequest)
			return
		}
		if len(g[0]) != 2 {
			p.l.Println("invalid URI")
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "invalid URI", http.StatusBadRequest)
			return
		}
		//calling PUT
		p.UpdateProducts(id, rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Product")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "unable to marshal json", http.StatusInternalServerError)
	}
}

func (p *Products) addProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST product")
	prod := &data.Product{}
	err := prod.GetJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusBadRequest)
	}

	p.l.Printf("Prod: %#v", prod)

}

func (p *Products) UpdateProducts(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT product")
	prod := &data.Product{}
	err := prod.GetJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusBadRequest)
	}
	err = data.UpdateProducts(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

}

//the logic follows
//request comes in thru newserver mux and handlers are activated, goes to PH, that is product
//handler, then triggers serveHTTO function which has a logic that tells whether its a
//PUT OR a GET command

//we had to write so many lines of code to get the ID from the URI. now this is where frameworks come into picture
