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
		return fmt.Errorf("Service error: %w", err)
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

