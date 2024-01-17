package grpc

import (
	"context"
	"log"
	"net"
	pb "sql_storage_layer/pkg/grpc/proto"
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

	employee, err := svc.service.GetByIDService(ctx, employeeID)
	if err != nil {
		return &pb.EmployeeResponse{}, err
	}

	response := &pb.EmployeeResponse{
		EmployeeId: employee.Employee_id,
		FirstName:  employee.First_name,
		LastName:   employee.Last_name,
	}

	return response, nil
}

func (svc *EmployeeServiceServer) CreateEmployee(ctx context.Context, req *pb.CreateEmployeeRequest) (*pb.CreateEmployeeResponse, error) {

	return &pb.CreateEmployeeResponse{}, nil

}

func (svc *EmployeeServiceServer) UpdateEmployee(ctx context.Context, req *pb.UpdateEmployeeRequest) (*pb.UpdateEmployeeResponse, error) {
	return &pb.UpdateEmployeeResponse{}, nil

}

func (svc *EmployeeServiceServer) DeleteEmployee(ctx context.Context, req *pb.DeleteEmployeeRequest) (*pb.DeleteEmployeeResponse, error) {
	return &pb.DeleteEmployeeResponse{}, nil
}

func StartServer() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a gRPC server
	srv := grpc.NewServer()

	// Register the ExampleServiceServer with the server
	pb.RegisterEmployeeServiceServer(srv, &EmployeeServiceServer{})

	// Serve the gRPC server
	if err := srv.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
