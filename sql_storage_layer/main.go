package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "nivea100"
	dbname   = "postgres"
)

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Initializing the Database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initializing the Repository
	repo := NewRepository(db)

	// Initializing the Service
	svc := NewService(*repo)

	// Creating a Router
	router := gin.Default()

	employeeGroup := router.Group("/v1/employee")
	{
		employeeGroup.POST("/", CreateEmployeeHandler(svc)) // Добавить возвращение employeeID в месседже
		employeeGroup.GET("/:id", GetEmployeeHandler(svc))
		employeeGroup.PUT("/:id", UpdateEmployeeHandler(svc)) // Status 200, но records отсутствуют
		employeeGroup.DELETE("/:id", DeleteEmployeeHandler(svc))
	}

	router.Run("localhost:8080")
}
