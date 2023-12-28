package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

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

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	parentCtx := context.Background()
	ctx, cancel := context.WithTimeout(parentCtx, 5*time.Second)
	defer cancel()

	repo := &Repository{DB: db}

	// empID, posID, err := repo.Create(ctx, employeeFixture, positionFixture)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("Created employee with id: %v and position with id %v \n", empID, posID)

	result, err := repo.GetByID(ctx, 7)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Founded employee with id: %v, firstname: %v, lastname: %v, position_id: %v",
		result.Employee_id, result.First_name, result.Last_name, result.Position_id)
}
