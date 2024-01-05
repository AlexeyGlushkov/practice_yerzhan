package main

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepository struct {
	Collection *mongo.Collection
}

type PositionRepository struct {
	Collection *mongo.Collection
}

func (er *EmployeeRepository) Create(ctx context.Context, firstName, lastName string) (string, error) {
	result, err := er.Collection.InsertOne(ctx, Employee{
		FirstName: firstName,
		LastName:  lastName,
	})
	if err != nil {
		return "", err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", errors.New("failed to get inserted ID")
	}

	insertedIDString := insertedID.Hex()

	return insertedIDString, nil
}

func (pr *PositionRepository) Create(ctx context.Context, employeeID, positionName string, salary int) error {
	_, err := pr.Collection.InsertOne(ctx, Position{
		EmployeeID:   employeeID,
		PositionName: positionName,
		Salary:       salary,
	})
	return err
}
