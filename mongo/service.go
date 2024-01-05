package main

import (
	"context"

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

	employeeID, err := s.EmployeeRepo.Create(ctx, firstName, lastName)
	if err != nil {
		session.AbortTransaction(ctx)
		return "", err
	}

	err = s.PositionRepo.Create(ctx, employeeID, positionName, salary)
	if err != nil {
		session.AbortTransaction(ctx)
		return "", err
	}

	err = session.CommitTransaction(ctx)
	if err != nil {
		return "", err
	}

	return employeeID, nil
}
