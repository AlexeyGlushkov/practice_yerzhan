package main

type CreateEmployeePayload struct {
	Employee Employee `json:"employee"`
	Position Position `json:"position"`
}

type GetEmployeeResponse struct {
	Employee Employee `json:"employee"`
}

type UpdateEmployeePayload struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
