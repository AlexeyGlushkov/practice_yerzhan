package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type Storage interface {
	Create(data interface{}) error
	Read(id int) (interface{}, error)
	Update(id int, newData interface{}) error
	Delete(id int) error
}

type Employee struct {
	Employee_id int
	First_name  string
	Last_name   string
	Position_id string
}

type EmployeeRepository struct {
	DB *sql.DB
}

type Position struct {
	Position_id   string
	Position_name string
	Salary        int
}

type PositionRepository struct {
	DB *sql.DB
}

func (er *EmployeeRepository) Create(employee Employee) error {
	_, err := er.DB.Exec("INSERT INTO employee (employee_id, first_name, last_name, position_id) VALUES ($1, $2, $3, $4)",
		employee.Employee_id, employee.First_name, employee.Last_name, employee.Position_id)

	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
