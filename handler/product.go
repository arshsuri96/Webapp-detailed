package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/arshsuri96/test/product-api/data"
	"github.com/gorilla/mux"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
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

func (p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {

	//mux handles id using vars map
	v := mux.Vars(r)
	id, err := strconv.Atoi(v["id"])
	if err != nil {
		http.Error(rw, "unable to convert id", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT product", id)

	prod := &data.Product{}
	err = prod.GetJSON(r.Body)
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
