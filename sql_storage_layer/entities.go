package main

type Employee struct {
	Employee_id string `json:"employee_id"`
	First_name  string `json:"first_name"`
	Last_name   string `json:"last_name"`
}

type Position struct {
	Position_id   string `json:"position_id"`
	Position_name string `json:"position_name"`
	Salary        int    `json:"salary"`
	Employee_id   string `json:"employee_id"`
}
