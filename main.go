package main

import (
	"fmt"
	"net/http"

	"github.com/MET-DEV/api-project/config"
	"github.com/MET-DEV/api-project/handlers"
	"github.com/MET-DEV/api-project/migrations"
)

func main() {
	fmt.Println("Server started")
	config.DbConfiguration()
	migrations.IndexMigration()

	server := &http.Server{
		Addr:    ":8080",
		Handler: handlers.IndexRouting(),
	}
	server.ListenAndServe()

}
