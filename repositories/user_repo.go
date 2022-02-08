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
func GetByIdUser(id int) models.User {
	var user models.User
	config.DB.First(&user, id)
	return user
}
func AddUser(user models.User) models.User {
	config.DB.Create(&user)
	return user
}
func UpdateUser(user models.User) models.User {
	config.DB.Save(&user)
	return user
}
func DeleteUser(id int) models.User {
	var user models.User
	config.DB.Where("id=?", id).Delete(&user)
	return user
}
