package main

type Employee struct {
	Employee_id string
	First_name  string
	Last_name   string
}

type Position struct {
	Position_id   string
	Position_name string
	Salary        int
	Employee_id   string
}
