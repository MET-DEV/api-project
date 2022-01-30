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
	//? <--------------------Category-------------------------->
	r.HandleFunc("/api/categories", GetAllCategory).Methods("GET")
	r.HandleFunc("/api/categories/{id}", GetCategoryById).Methods("GET")
	r.HandleFunc("/api/categories", AddCategory).Methods("POST")
	r.HandleFunc("/api/categories/{id}", DeleteCategory).Methods("DELETE")
	r.HandleFunc("/api/categories/update", UpdateCategory).Methods("PATCH")

	return r
}
