package transport

import (
	"context"
	"fmt"
	pb "github.com/paulnasdaq/fms-v2/common/vehicles"
	"github.com/paulnasdaq/fms-v2/vehicles/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
)

type GRPCServer struct {
	pb.UnimplementedVehiclesServiceServer
	vehicleService service.Service
}

func NewGRPCServer(s *service.Service) *GRPCServer {
	return &GRPCServer{vehicleService: s, UnimplementedVehiclesServiceServer: pb.UnimplementedVehiclesServiceServer{}}
}
func (s *GRPCServer) AddVehicleModel(ctx context.Context, r *pb.AddVehicleModelRequest) (*pb.AddVehicleModelResponse, error) {
	res, err := s.vehicleService.AddVehicleModel(r.Name)
	if err != nil {
		return nil, err
	}
	return &pb.AddVehicleModelResponse{Name: res.Name, Id: res.ID}, err
}
func (s *GRPCServer) AddVehicle(ctx context.Context, r *pb.AddVehicleRequest) (*pb.AddVehicleResponse, error) {
	res, err := s.vehicleService.AddVehicle(r.RegistrationNumber, r.ChassisNumber, r.ModelId)
	if err != nil {
		return nil, err
	}
	return &pb.AddVehicleResponse{Error: nil, ModelId: res.Model.ID, Id: res.ID, ChassisNumber: res.ChassisNumber, RegistrationNumber: res.RegNumber}, err
}
func (s *GRPCServer) Listen(port int) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}

	server := grpc.NewServer()
	pb.RegisterVehiclesServiceServer(server, s)
	reflection.Register(server)
	return server.Serve(lis)
}
