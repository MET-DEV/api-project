package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/MET-DEV/api-project/models"
	"github.com/MET-DEV/api-project/repositories"
	"github.com/gorilla/mux"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repositories.GetAllProducts())
}
func AddProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&product)
	var productDb models.Product = repositories.AddProduct(product)
	json.NewEncoder(w).Encode(productDb)
}
func GetProductById(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	w.Header().Set("Content-Type", "application-json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Print(err)
	}
	product = repositories.GetByIdProduct(id)
	json.NewEncoder(w).Encode(product)
}
func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
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
	json.NewEncoder(w).Encode(productDb)
}
