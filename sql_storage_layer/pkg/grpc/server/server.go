package grpc

import (
	"context"
	prt "sql_storage_layer/pkg/grpc/proto"
	servc "sql_storage_layer/pkg/service"
)

type EmployeeServiceServer struct {
	service *servc.Service
}

func (svc *EmployeeServiceServer) GetEmployeeByID(ctx context.Context, req *prt.EmployeeRequest) (*prt.EmployeeResponse, error) {

	employeeID := req.EmployeeId

	employee, err := svc.service.GetByIDService(ctx, employeeID)
	if err != nil {
		return &prt.EmployeeResponse{}, err
	}

	response := &prt.EmployeeResponse{
		EmployeeId: employee.Employee_id,
		FirstName:  employee.First_name,
		LastName:   employee.Last_name,
	}

	return response, nil
}
