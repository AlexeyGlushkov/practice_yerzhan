package main

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	parentCtx := context.Background()
	ctx, cancel := context.WithTimeout(parentCtx, 5*time.Second)
	defer cancel()

	// Connection to MongoDB
	uri := "mongodb+srv://kousetsu:kousetsuxo@kousetsu.krfgpxk.mongodb.net/?retryWrites=true&w=majority"

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	db := client.Database("practice")

	// Collections
	employeeCollection := db.Collection("employees")
	positionCollection := db.Collection("positions")

	// Repo's
	employeeRepo := &EmployeeRepository{Collection: employeeCollection}
	positionRepo := &PositionRepository{Collection: positionCollection}

	// Service
	svc := &Service{
		EmployeeRepo: employeeRepo,
		PositionRepo: positionRepo,
		Database:     db,
	}

	// Service usage
	employeeID, err := svc.Create(ctx, "John", "Doe", "Manager", 5000)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Created Employee with ID:", employeeID)
}
