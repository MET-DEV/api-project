package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName string   `json:"product_name" validate:"required" `
	Description string   `json:"description" validate:"required"`
	Price       float32  `json:"price" validate:"required"`
	CategoryID  uint     `json:"category_id" validate:"required"`
	Category    Category `gorm:"association_foreignkey:CategoryID" json:"category"`
}
