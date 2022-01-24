package main

import (
	"fmt"

	"github.com/MET-DEV/api-project/config"
	"github.com/MET-DEV/api-project/migrations"
	"github.com/MET-DEV/api-project/repositories"
)

func main() {
	fmt.Println("Server started")
	config.DbConfiguration()
	migrations.ProductMigration()
	fmt.Println(repositories.GetAll())

}
