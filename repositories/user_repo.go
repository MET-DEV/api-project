package repositories

import (
	"github.com/MET-DEV/api-project/config"
	"github.com/MET-DEV/api-project/models"
)

func GetAllUsers() []models.User {
	var users []models.User
	config.DB.Find(&users)
	return users
}
func AddUser(user models.User) models.User {
	config.DB.Create(&user)
	return user
}
