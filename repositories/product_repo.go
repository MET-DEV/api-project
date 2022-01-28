package repositories

import (
	"github.com/MET-DEV/api-project/config"
	"github.com/MET-DEV/api-project/models"
)

func GetAllProducts() []models.Product {
	var products []models.Product
	config.DB.Find(&products)
	return products
}
func GetByIdProduct(id int) models.Product {
	var product models.Product
	config.DB.First(&product, id)
	return product
}
func AddProduct(product models.Product) models.Product {
	config.DB.Create(&product)
	return product

}
func DeleteProduct(id int) models.Product {
	var product models.Product
	config.DB.Where("id=?", id).Delete(&product)
	return product
}
func UpdateProduct(product models.Product) models.Product {
	config.DB.Save(&product)
	return product

}
