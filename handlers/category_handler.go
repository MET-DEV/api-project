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

func GetAllCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(repositories.GetAllCategory())
}
func GetCategoryById(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
	}
	category = repositories.GetByIdCategory(id)
	json.NewEncoder(w).Encode(category)
}
func AddCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&category)
	var categoryDb models.Category = repositories.AddCategory(category)
	json.NewEncoder(w).Encode(categoryDb)
}
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Println(err)
	}
	repositories.DeleteCategory(id)
	json.NewEncoder(w).Encode("Deleted")
}
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var category models.Category
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&category)
	var categoryDb models.Category = repositories.UpdateCategory(category)
	json.NewEncoder(w).Encode(categoryDb)
}
