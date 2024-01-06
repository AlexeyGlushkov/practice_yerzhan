package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Service struct {
	EmployeeRepo *EmployeeRepository
	PositionRepo *PositionRepository
	Database     *mongo.Database
}

func (s *Service) CreateService(ctx context.Context, firstName, lastName, positionName string, salary int) (string, error) {
	session, err := s.Database.Client().StartSession()
	if err != nil {
		return "", err
	}
	defer session.EndSession(ctx)

	err = session.StartTransaction()
	if err != nil {
		return "", err
	}

	employeeIDstring, err := s.EmployeeRepo.Create(ctx, firstName, lastName)
	if err != nil {
		session.AbortTransaction(ctx)
		return "", err
	}

	employeeID, err := primitive.ObjectIDFromHex(employeeIDstring)
	if err != nil {
		session.AbortTransaction(ctx)
		return "", err
	}

	err = s.PositionRepo.Create(ctx, positionName, salary, employeeID)
	if err != nil {
		session.AbortTransaction(ctx)
		return "", err
	}

	err = session.CommitTransaction(ctx)
	if err != nil {
		return "", err
	}

	return employeeIDstring, nil
}
