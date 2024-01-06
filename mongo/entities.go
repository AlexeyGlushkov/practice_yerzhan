package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	EmployeeID primitive.ObjectID `bson:"employee_id,omitempty"`
	FirstName  string             `bson:"first_name"`
	LastName   string             `bson:"last_name"`
}

type Position struct {
	PositionID   primitive.ObjectID `bson:"position_id,omitempty"`
	PositionName string             `bson:"position_name"`
	Salary       int                `bson:"salary"`
	EmployeeID   primitive.ObjectID `bson:"employee_id,omitempty"`
}
