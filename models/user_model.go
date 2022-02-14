package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName   string `json:"first_name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Email       string `json:"email" validate:"required"`
	Password    string `json:"password" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	BirthYear   int    `json:"birth_date" validate:"required"`
}
