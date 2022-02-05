package migrations

import (
	"github.com/MET-DEV/api-project/config"
	"github.com/MET-DEV/api-project/models"
)

func UserMigration() {
	config.DB.AutoMigrate(&models.User{})
}
