package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type EmployeeRepository struct {
	DB *sql.DB
}

type PositionRepository struct {
	DB *sql.DB
}

func (er *EmployeeRepository) CreateEmployee(employee Employee) error {

	_, err := er.DB.Exec("INSERT INTO employee (employee_id, first_name, last_name, position_id)"+
		"VALUES ($1, $2, $3, $4)",
		employee.Employee_id, employee.First_name, employee.Last_name, employee.Position_id)

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func (pr *PositionRepository) CreatePosition(ctx context.Context, position Position) (positionID string, err error) {

	fail := func(err error) (string, error) {
		return "", fmt.Errorf("CreatePosition: %v", err)
	}

	tx, err := pr.DB.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}
	defer tx.Rollback()

	var position_id string
	if err = tx.QueryRowContext(ctx, "INSERT INTO position (position_name, salary)"+
		" VALUES ($1, $2) RETURNING position_id", position.Position_name, position.Salary).Scan(&position_id); err != nil {
		if err == sql.ErrNoRows {
			return fail(fmt.Errorf("no rows returned: %v", err))
		}
		return fail(fmt.Errorf("error inserting position: %v", err))
	}

	if err = tx.Commit(); err != nil {
		return fail(err)
	}

	return position_id, nil

}
