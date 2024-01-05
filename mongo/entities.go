package main

type Employee struct {
	EmployeeID string `bson:"employee_id"`
	FirstName  string `bson:"first_name"`
	LastName   string `bson:"last_name"`
}

type Position struct {
	PositionID   string `bson:"position_id"`
	PositionName string `bson:"position_name"`
	Salary       int    `bson:"salary"`
	EmployeeID   string `bson:"employee_id"`
}
