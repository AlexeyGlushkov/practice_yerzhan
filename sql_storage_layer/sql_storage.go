package main

import (
	"database/sql"
	"fmt"
	"log"

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

	empRepo := &EmployeeRepository{DB: db}

	testEmployee := Employee{
		Employee_id: 5,
		First_name:  "Jack",
		Last_name:   "Suuui",
		Position_id: "1b76ab62-954c-492c-9877-db4f24881376",
	}

	err = empRepo.Create(testEmployee)
	if err != nil {
		fmt.Printf("Failed to create employee, err: %v \n", err)
		return
	}

	fmt.Println("Employee created successfully")
}
