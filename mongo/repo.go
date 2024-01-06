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

	objectID := primitive.NewObjectID()

	employee := Employee{
		EmployeeID: objectID,
		FirstName:  firstName,
		LastName:   lastName,
	}

	result, err := er.Collection.InsertOne(ctx, employee)
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

func (pr *PositionRepository) Create(ctx context.Context, positionName string, salary int, employeeID primitive.ObjectID) error {

	objectID := primitive.NewObjectID()

	position := Position{
		PositionID:   objectID,
		PositionName: positionName,
		Salary:       salary,
		EmployeeID:   employeeID,
	}

	_, err := pr.Collection.InsertOne(ctx, position)
	return err
}
