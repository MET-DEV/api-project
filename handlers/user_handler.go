package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/MET-DEV/api-project/models"
	"github.com/MET-DEV/api-project/repositories"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories.GetAllUsers())

}
func AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&user)
	var userDb models.User = repositories.AddUser(user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userDb)

}
