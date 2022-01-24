package repositories

import (
	"github.com/MET-DEV/api-project/config"
	"github.com/MET-DEV/api-project/models"
)

func GetAll() []models.Product {
	var products []models.Product
	config.DB.Find(&products)
	return products
}
