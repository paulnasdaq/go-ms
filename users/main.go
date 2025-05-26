package main

import (
	"context"
	"github.com/google/uuid"
	pb "github.com/paulnasdaq/fms-v2/common/users"
	"github.com/paulnasdaq/fms-v2/users/db"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type GRPCServer struct {
	pb.UnimplementedUsersServiceServer
	r db.Repository
}

func (s *GRPCServer) AddUser(ctx context.Context, r *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := s.r.Users().Add(r.Email)
	if err != nil {
		log.Println("Failed to create user", err)
		return nil, err
	}
	return &pb.CreateUserResponse{Id: user.ID.String(), Email: user.Email}, nil
}

func (s *GRPCServer) GetUSer(ctx context.Context, r *pb.GetUserRequest) (*pb.CreateUserResponse, error) {
	user, err := s.r.Users().Get(uuid.MustParse(r.Id))
	if err != nil {
		log.Println("Failed to get user", err)
		return nil, err
	}
	return &pb.CreateUserResponse{Id: user.ID.String(), Email: user.Email}, nil
}
func (s *GRPCServer) GetByEmail(ctx context.Context, r *pb.GetByEmailRequest) (*pb.CreateUserResponse, error) {
	user, err := s.r.Users().GetByEmail(r.Email)
	if err != nil {
		log.Println("Failed to find user", err)
		return nil, err
	}
	return &pb.CreateUserResponse{Id: user.ID.String(), Email: user.Email}, nil
}

func main() {
	repo, err := db.NewRepository()
	if err != nil {
		log.Fatal("Failed to start service", err)
	}
	s := &GRPCServer{r: repo}

	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal("Failed to start service", err)
	}

	server := grpc.NewServer()
	pb.RegisterUsersServiceServer(server, s)
	reflection.Register(server)

	log.Fatal(server.Serve(lis))
}
