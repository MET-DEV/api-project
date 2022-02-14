package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/MET-DEV/api-project/models"
	"github.com/MET-DEV/api-project/repositories"
	"github.com/MET-DEV/api-project/security"
	"github.com/MET-DEV/api-project/validations"
	"github.com/gorilla/mux"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(repositories.GetAllUsers())

}
func GetUserById(w http.ResponseWriter, r *http.Request) {
	var user models.User
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Id hatalı")
		return

	}
	w.WriteHeader(http.StatusOK)
	user = repositories.GetByIdUser(id)
	json.NewEncoder(w).Encode(user)
}
func AddUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&user)
	e := validations.UserValidator(user)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(e.Field() + " is required")
		return
	}
	user.Password = security.HashPassword(user.Password)
	if user.Password == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Parola şifrelenirken hata oluştu")
		return
	}

	var userDb models.User = repositories.AddUser(user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userDb)

}
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application-json")
	var user models.User
	w.Header().Set("Content-Type", "application/json")
	json.NewDecoder(r.Body).Decode(&user)
	var userDb models.User = repositories.UpdateUser(user)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(userDb)
}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Id hatalı")
		return
	}
	w.WriteHeader(http.StatusOK)
	repositories.DeleteUser(id)
	json.NewEncoder(w).Encode("Deleted")
}
