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

func GetAllCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories.GetAllCategory())
}
func GetCategoryById(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Id hatalı")
		return

	}
	w.WriteHeader(http.StatusOK)
	category = repositories.GetByIdCategory(id)
	json.NewEncoder(w).Encode(category)
}
func AddCategory(w http.ResponseWriter, r *http.Request) {
	var category models.Category
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&category)
	e := validations.CategoryValidator(category)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e.Field() + " is required")
		return
	}
	w.WriteHeader(http.StatusOK)
	var categoryDb models.Category = repositories.AddCategory(category)
	json.NewEncoder(w).Encode(categoryDb)
}
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Id hatalı")
		return
	}
	w.WriteHeader(http.StatusOK)
	repositories.DeleteCategory(id)
	json.NewEncoder(w).Encode("Deleted")
}
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var category models.Category
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&category)
	var categoryDb models.Category = repositories.UpdateCategory(category)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(categoryDb)
}
