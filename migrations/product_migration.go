package migrations

import (
	"github.com/MET-DEV/api-project/config"
	"github.com/MET-DEV/api-project/models"
)

func ProductMigration() {
	config.DB.AutoMigrate(&models.Product{})
}
func CategoryMigration() {
	config.DB.AutoMigrate(&models.Category{})
}
