package handlers

import (
	"github.com/gorilla/mux"
)

func IndexRouting() *mux.Router {
	r := mux.NewRouter()
	//? <--------------------Products-------------------------->
	r.HandleFunc("/api/products", GetAllProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", GetProductById).Methods("GET")
	r.HandleFunc("/api/products", AddProduct).Methods("POST")
	r.HandleFunc("/api/products/{id}", DeleteProduct).Methods("DELETE")
	r.HandleFunc("/api/products/update", UpdateProduct).Methods("PATCH")

	return r
}
