package grpc

import (
	"context"
	"fmt"
	"log"
	"net"
	pb "sql_storage_layer/pkg/grpc/proto"
	"sql_storage_layer/pkg/models"
	servc "sql_storage_layer/pkg/service"

	"google.golang.org/grpc"
)

type EmployeeServiceServer struct {
	service *servc.Service
	pb.UnsafeEmployeeServiceServer
}

// mustEmbedUnimplementedEmployeeServiceServer implements emppos_proto.EmployeeServiceServer.
func (*EmployeeServiceServer) mustEmbedUnimplementedEmployeeServiceServer() {
	panic("unimplemented")
}

func (svc *EmployeeServiceServer) GetEmployeeByID(ctx context.Context, req *pb.EmployeeRequest) (*pb.EmployeeResponse, error) {

	employeeID := req.EmployeeId
	log.Printf("GetEmployeeByID called with EmployeeID: %s", employeeID)

	employee, err := svc.service.GetByIDService(ctx, employeeID)
	if err != nil {
		log.Printf("Error in GetByIDService: %v", err)
		return &pb.EmployeeResponse{}, err
	}

	response := &pb.EmployeeResponse{
		EmployeeId: employee.Employee_id,
		FirstName:  employee.First_name,
		LastName:   employee.Last_name,
	}

	log.Printf("GetEmployeeByID successfully completed")
	return response, nil
}

func (svc *EmployeeServiceServer) CreateEmployee(ctx context.Context, req *pb.CreateEmployeeRequest) (*pb.CreateEmployeeResponse, error) {

	reqEmployee := models.Employee{
		First_name: req.FirstName,
		Last_name:  req.LastName,
	}

	reqPosition := models.Position{
		Position_name: req.PositionName,
		Salary:        int(req.Salary),
	}

	err := svc.service.CreateService(ctx, reqEmployee, reqPosition)
	if err != nil {
		log.Printf("Error in CreateService: %v", err)
		return nil, err
	}

	log.Printf("Employee successfully created")
	return &pb.CreateEmployeeResponse{}, nil

}

func (svc *EmployeeServiceServer) UpdateEmployee(ctx context.Context, req *pb.UpdateEmployeeRequest) (*pb.UpdateEmployeeResponse, error) {

	employeeID := req.EmployeeId
	firstName := req.FirstName
	lastName := req.LastName

	log.Printf("UpdateEmployee called with EmployeeID: %s", employeeID)

	err := svc.service.UpdateEmployeeService(ctx, employeeID, firstName, lastName)
	if err != nil {
		log.Printf("Error in UpdateEmployeeService: %v", err)
		return nil, err
	}

	log.Printf("Employee successfully updated")
	return &pb.UpdateEmployeeResponse{}, nil
}

func (svc *EmployeeServiceServer) DeleteEmployee(ctx context.Context, req *pb.DeleteEmployeeRequest) (*pb.DeleteEmployeeResponse, error) {
	employeeID := req.EmployeeId
	log.Printf("DeleteEmployee called with EmployeeID: %s", employeeID)

	err := svc.service.DeleteService(ctx, employeeID)
	if err != nil {
		log.Printf("Error in DeleteService: %v", err)
		return nil, err
	}

	log.Printf("Employee successfully deleted")
	return &pb.DeleteEmployeeResponse{}, nil
}

func StartServer(service *servc.Service) {
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Create a gRPC server
	srv := grpc.NewServer()

	// Register the ExampleServiceServer with the server
	pb.RegisterEmployeeServiceServer(srv, &EmployeeServiceServer{
		service: service,
	})

	log.Println("gRPC server is starting...")

	// Serve the gRPC server
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
