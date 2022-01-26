package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MET-DEV/api-project/repositories"
)

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	repositories.GetAll()
	json.NewEncoder(w).Encode(repositories.GetAll())
}
