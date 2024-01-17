package http

import (
	"sql_storage_layer/pkg/models"
)

type CreateEmployeePayload struct {
	Employee models.Employee `json:"employee"`
	Position models.Position `json:"position"`
}

type GetEmployeeResponse struct {
	Employee models.Employee `json:"employee"`
}

type UpdateEmployeePayload struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
