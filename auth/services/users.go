package services

import (
	"context"
	"github.com/paulnasdaq/fms-v2/auth/db"
	"github.com/paulnasdaq/fms-v2/common"
	users_pb "github.com/paulnasdaq/fms-v2/common/users"
	"log"
)

type AuthUsersService struct {
	repository db.AuthRepository
	clients    *common.ClientManager
	userPass   *AuthUserPassService
}
type User struct {
	ID    string
	Email string
}

func NewAuthUsersService(r db.AuthRepository, clients *common.ClientManager, userPassService *AuthUserPassService) *AuthUsersService {
	return &AuthUsersService{repository: r, clients: clients, userPass: userPassService}
}
func (s *AuthUsersService) Add(email string, password string) (*User, error) {
	usersService, err := s.clients.Users()
	if err != nil {
		log.Println("Failed to get users service", err)
		return nil, err
	}
	ctx := context.Context(context.Background())
	newUser, err := usersService.AddUser(ctx, &users_pb.CreateUserRequest{
		Email: email,
	})
	if err != nil {
		log.Println("Failed to create user", err)
		return nil, err
	}

	err = s.userPass.Add(newUser.Id, password)
	if err != nil {
		log.Println("Failed to add user", err)
		return nil, err
	}
	return &User{Email: newUser.Email, ID: newUser.Id}, nil
}
