package main

import (
	"context"
	"fmt"
)

type Service struct {
	Repo Repository
}

func (svc *Service) CreateService(ctx context.Context, emp Employee, pos Position) error {

	fail := func(err error) error {
		return fmt.Errorf("Create Service error: %w", err)
	}

	tx, err := svc.Repo.DB.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}
	defer tx.Rollback()

	employeeID, err := svc.Repo.CreateEmployee(ctx, tx, emp.First_name, emp.Last_name)
	if err != nil {
		return fail(err)
	}

	if err = svc.Repo.CreatePosition(ctx, tx, employeeID, pos.Position_name, pos.Salary); err != nil {
		return fail(err)
	}

	if err = tx.Commit(); err != nil {
		return fail(err)
	}

	return nil
}

func (svc *Service) GetByIDService(ctx context.Context, employeeID string) (Employee, error) {

	resEmp, err := svc.Repo.GetByID(ctx, employeeID)
	if err != nil {
		return Employee{}, fmt.Errorf("GetByID Service error: %w", err)
	}

	return resEmp, nil
}

func (svc *Service) UpdateEmployeeService(ctx context.Context, empID, fName, lName string) (Employee, error) {

	resEmp, err := svc.Repo.UpdateEmployee(ctx, empID, fName, lName)
	if err != nil {
		return Employee{}, fmt.Errorf("UpdateEmployee Service error: %w", err)
	}

	return resEmp, nil
}

func (svc *Service) UpdatePositionService(ctx context.Context, posID, posName string, salary int) (Position, error) {

	resPos, err := svc.Repo.UpdatePosition(ctx, posID, posName, salary)
	if err != nil {
		return Position{}, fmt.Errorf("UpdatePosition Service error: %w", err)
	}

	return resPos, nil
}

func (svc *Service) DeleteService(ctx context.Context, employeeID string) error {

	fail := func(err error) error {
		return fmt.Errorf("Delete Service error: %w", err)
	}

	tx, err := svc.Repo.DB.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}
	defer tx.Rollback()

	if err = svc.Repo.Delete(ctx, tx, employeeID); err != nil {
		return fail(err)
	}

	if err = tx.Commit(); err != nil {
		return fail(err)
	}

	return nil
}
