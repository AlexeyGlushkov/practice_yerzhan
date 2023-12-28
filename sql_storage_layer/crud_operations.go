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

func (r *Repository) Create(ctx context.Context, emp Employee, pos Position) (int, string, error) {

	fail := func(err error) (int, string, error) {
		return 0, "", fmt.Errorf("Create: %w", err)
	}

	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return fail(err)
	}
	defer tx.Rollback()

	insertPositionQuery := `INSERT INTO position (position_name, salary)
	VALUES ($1, $2) RETURNING position_id;`

	var positionID string

	err = tx.QueryRowContext(ctx, insertPositionQuery, pos.Position_name, pos.Salary).Scan(&positionID)
	if err != nil {
		return fail(err)
	}

	insertEmployeeQuery := `
	INSERT INTO employee (first_name, last_name, position_id)
	VALUES ($1, $2, $3) RETURNING employee_id;`

	var employeeID int

	err = tx.QueryRowContext(ctx, insertEmployeeQuery, emp.First_name, emp.Last_name, positionID).Scan(&employeeID)
	if err != nil {
		return fail(err)
	}

	if err = tx.Commit(); err != nil {
		return fail(err)
	}

	return employeeID, positionID, nil
}

func (r *Repository) GetByID(ctx context.Context, empID int) (Employee, error) {

	fail := func(err error) (Employee, error) {
		return Employee{}, fmt.Errorf("GetByID: %w", err)
	}

	selectStatement := `
	SELECT employee_id, first_name, last_name, position_id FROM employee
	WHERE employee_id = $1;`

	var resEmp Employee

	row := r.DB.QueryRowContext(ctx, selectStatement, empID)
	err := row.Scan(&resEmp.Employee_id, &resEmp.First_name, &resEmp.Last_name, &resEmp.Position_id)

	if err != nil {
		return fail(err)
	}

	return resEmp, nil
}

func (r *Repository) UpdateEmployee(ctx context.Context, empID int, fName, lName, posID string) (Employee, error) {

	fail := func(err error) (Employee, error) {
		return Employee{}, fmt.Errorf("UpdateEmployee: %w", err)
	}

	updateStatement := `
	UPDATE employee
	SET first_name = $2, last_name = $3, position_id = $4
	WHERE employee_id = $1
	RETURNING employee_id, first_name, last_name, position_id;`

	var updatedEmp Employee

	row := r.DB.QueryRowContext(ctx, updateStatement, empID, fName, lName, posID)
	err := row.Scan(&updatedEmp.Employee_id, &updatedEmp.First_name, &updatedEmp.Last_name, &updatedEmp.Position_id)

	if err != nil {
		return fail(err)
	}

	return updatedEmp, nil
}

func (r *Repository) UpdatePosition(ctx context.Context, posID, posName string, salary int) (Position, error) {

	fail := func(err error) (Position, error) {
		return Position{}, fmt.Errorf("UpdateEmployee: %w", err)
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
