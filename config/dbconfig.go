package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

const db_string = "host=localhost user=postgres password=12345 dbname=ETradeDB port=5432 sslmode=disable"

func DbConfiguration() {
	DB, err = gorm.Open(postgres.Open(db_string), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
	}

}
