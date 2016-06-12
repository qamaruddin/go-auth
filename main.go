package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/qamaruddin/auth/models"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/", NotImplemented).Methods("GET")
	r.Handle("/products", models.ProductsHandler)

	http.ListenAndServe(":3000", r)
}

//Handle not Implemented routes by "NotImplemented ..."
var NotImplemented = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not Implemented"))
})
