package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MET-DEV/api-project/models"
	"github.com/MET-DEV/api-project/repositories"
	"github.com/MET-DEV/api-project/validations"
	"github.com/gorilla/mux"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories.GetAllProducts())
}
func AddProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&product)
	validateProduct := product
	validateProduct.Category.CategoryName = "passthisfield"
	e := validations.ProductValidator(validateProduct)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e.Field() + " is required")
		return
	}
	var productDb models.Product = repositories.AddProduct(product)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productDb)
}

func GetProductById(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	w.Header().Set("Content-Type", "application-json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Id hatalı")
		return
	}
	product = repositories.GetByIdProduct(id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Id hatalı")
		return
	}
	repositories.DeleteProduct(id)
	json.NewEncoder(w).Encode("Deleted")
}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var product models.Product
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&product)
	var productDb models.Product = repositories.UpdateProduct(product)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(productDb)
}
