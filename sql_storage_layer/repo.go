package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) CreateEmployee(ctx context.Context, tx *sql.Tx, empFirstname, empLastname string) (string, error) {
	insertEmployeeQuery := `INSERT INTO employee (first_name, last_name)
	VALUES ($1, $2) RETURNING employee_id;`

	var employeeID string

	err := tx.QueryRowContext(ctx, insertEmployeeQuery, empFirstname, empLastname).Scan(&employeeID)
	if err != nil {
		return "", err
	}

	return employeeID, nil
}

func (r *Repository) CreatePosition(ctx context.Context, tx *sql.Tx, empID, posName string, salary int) error {
	insertPositionQuery := `INSERT INTO position (employee_id, position_name, salary)
	VALUES ($1, $2, $3)`

	_, err := tx.ExecContext(ctx, insertPositionQuery, empID, posName, salary)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) GetByID(ctx context.Context, empID string) (Employee, error) {

	fail := func(err error) (Employee, error) {
		return Employee{}, fmt.Errorf("GetByID error: %w", err)
	}

	selectStatement := `
	SELECT employee_id, first_name, last_name FROM employee
	WHERE employee_id = $1;`

	var resEmp Employee

	row := r.DB.QueryRowContext(ctx, selectStatement, empID)
	err := row.Scan(&resEmp.Employee_id, &resEmp.First_name, &resEmp.Last_name)

	if err != nil {
		return fail(err)
	}

	return resEmp, nil
}

func (r *Repository) UpdateEmployee(ctx context.Context, empID, fName, lName string) (Employee, error) {

	fail := func(err error) (Employee, error) {
		return Employee{}, fmt.Errorf("UpdateEmployee error: %w", err)
	}

	updateStatement := `
	UPDATE employee
	SET first_name = $2, last_name = $3
	WHERE employee_id = $1
	RETURNING employee_id, first_name, last_name;`

	var updatedEmp Employee

	row := r.DB.QueryRowContext(ctx, updateStatement, empID, fName, lName)
	err := row.Scan(&updatedEmp.Employee_id, &updatedEmp.First_name, &updatedEmp.Last_name)

	if err != nil {
		return fail(err)
	}

	return updatedEmp, nil
}

func (r *Repository) UpdatePosition(ctx context.Context, posID, posName string, salary int) (Position, error) {

	fail := func(err error) (Position, error) {
		return Position{}, fmt.Errorf("UpdatePosition error : %w", err)
	}

	updateStatement := `
	UPDATE position
	SET position_name = $2, salary = $3
	WHERE position_id = $1
	RETURNING position_id, position_name, salary;`

	var updatedPos Position

	row := r.DB.QueryRowContext(ctx, updateStatement, posID, posName, salary)
	err := row.Scan(&updatedPos.Position_id, &updatedPos.Position_name, &updatedPos.Salary)

	if err != nil {
		return fail(err)
	}

	return updatedPos, nil
}

func (repo *Repository) Delete(ctx context.Context, tx *sql.Tx, employeeID string) error {

	fail := func(err error) error {
		return fmt.Errorf("Delte error: %w", err)
	}

	if _, err := tx.ExecContext(ctx, "DELETE FROM position WHERE employee_id = $1", employeeID); err != nil {
		return fail(err)
	}

	if _, err := tx.ExecContext(ctx, "DELETE FROM employee WHERE employee_id = $1", employeeID); err != nil {
		return fail(err)
	}

	return nil
}
