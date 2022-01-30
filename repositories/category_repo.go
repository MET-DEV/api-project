package repositories

import (
	"github.com/MET-DEV/api-project/config"
	"github.com/MET-DEV/api-project/models"
)

func GetAllCategory() []models.Category {
	var categories []models.Category
	config.DB.Find(&categories)
	return categories
}
func GetByIdCategory(id int) models.Category {
	var category models.Category
	config.DB.First(&category, id)
	return category
}
func AddCategory(category models.Category) models.Category {
	config.DB.Create(&category)
	return category

}
func UpdateCategory(category models.Category) models.Category {
	config.DB.Save(&category)
	return category

}
func DeleteCategory(id int) models.Category {
	var category models.Category
	config.DB.Where("id=?", id).Delete(&category)
	return category
}
