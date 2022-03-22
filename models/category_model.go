package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	CategoryName string    `json:"category_name" validate:"required"`
	Products     []Product `gorm:"foreignkey:CategoryID"`
}
