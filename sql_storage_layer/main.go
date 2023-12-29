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

	// SERVICE init

	svc := &Service{Repo: *repo}

	// svc.Create() -> empID
	err = svc.CreateService(ctx, employeeFixture, positionFixture)
	if err != nil {
		log.Fatal("CreateService failed")
	}

	fmt.Println("Successfully created")

	// svc.getByID(empID)
}
