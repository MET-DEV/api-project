package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ProductName string   `json:"product_name"`
	Description string   `json:"description"`
	Price       float32  `json:"price"`
	CategoryID  int      `json:"category_id"`
	Category    Category `json:"category"`
}
