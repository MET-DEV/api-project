package handlers

import (
	"github.com/gorilla/mux"
)

func IndexRouting() *mux.Router {
	r := mux.NewRouter()
	//? <--------------------Products-------------------------->
	r.HandleFunc("/api/products", GetAllProducts).Methods("GET")
	return r
}
