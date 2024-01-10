package main

import (
	"database/sql"
	"fmt"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "sql_storage_layer/docs"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "nivea100"
	dbname   = "postgres"
)

// @title Employee Service API
// @version 1.0
// @description Application for operations on employee and position
// @host localhost:8080
// @BasePath /v1
func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Connecting to the PostgreSQL database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initializing the Repository
	repo := NewRepository(db)

	// Initializing the Service
	svc := NewService(*repo)

	// Setting up the router
	router := SetupRouter(svc)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Starting the server
	router.Run("localhost:8080")
}
