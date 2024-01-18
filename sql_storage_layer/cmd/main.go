package main

import (
	"database/sql"
	"fmt"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "sql_storage_layer/docs"
	server "sql_storage_layer/pkg/grpc/server"
	http "sql_storage_layer/pkg/http"
	cache "sql_storage_layer/pkg/repo/cache"
	repo "sql_storage_layer/pkg/repo/postgres"
	service "sql_storage_layer/pkg/service"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "nivea100"
	dbname   = "postgres"
)

const (
	redisAddr     = "localhost:6379"
	redisPassword = ""
	redisDB       = 0
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
	repo := repo.NewRepository(db)

	// Initializing the Cache
	client, err := cache.NewRedisClient(redisAddr, redisPassword, redisDB)
	if err != nil {
		log.Fatal(err)
	}

	// Initializing the Service
	svc := service.NewService(*repo, *client)

	// Setting routes for the HTTP server
	router := http.SetupRouter(svc)

	// Setting routes for Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Starting the HTTP Server
	go func() {
		if err := router.Run("localhost:8080"); err != nil {
			log.Fatalf("Failed to run HTTP server: %v", err)
		}
	}()

	// Starting a gRPC server
	server.StartServer(svc)
}
