package main

import (
	"context"
	"github.com/paulnasdaq/fms-v2/auth/db"
	"github.com/paulnasdaq/fms-v2/auth/services"
	"github.com/paulnasdaq/fms-v2/common"
	pb "github.com/paulnasdaq/fms-v2/common/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type GRPCServer struct {
	pb.UnimplementedAuthServer
	usersService  *services.AuthUsersService
	tokensService services.AuthTokenService
}

func (s *GRPCServer) AddUser(ctx context.Context, r *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	user, err := s.usersService.Add(r.Email, r.Password)
	if err != nil {
		log.Println("Failed to create user", err)
		return nil, err
	}
	return &pb.AddUserResponse{Email: user.Email, Id: user.ID}, err
}

func (s *GRPCServer) GetToken(ctx context.Context, r *pb.GetTokenRequest) (*pb.GetTokenResponse, error) {
	token, err := s.tokensService.Create(r.Email, r.Password)
	if err != nil {
		return nil, err
	}
	return &pb.GetTokenResponse{Token: token}, nil
}
func main() {
	repository, err := db.NewRepository()
	if err != nil {
		log.Fatal(err)
	}
	clients := common.NewClientManager()
	userPassService := services.NewAuthUserPassService(repository)
	usersService := services.NewAuthUsersService(repository, clients, userPassService)

	tokensService := services.NewAuthTokenService(repository, clients, userPassService)

	server := &GRPCServer{usersService: usersService, tokensService: tokensService}

	lis, err := net.Listen("tcp", ":3003")
	if err != nil {
		log.Fatal("Failed to start service", err)
	}

	s := grpc.NewServer()
	pb.RegisterAuthServer(s, server)
	reflection.Register(s)
	log.Fatal(s.Serve(lis))
}
