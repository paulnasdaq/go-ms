package services

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/paulnasdaq/fms-v2/auth/db"
	"github.com/paulnasdaq/fms-v2/common"
	pb "github.com/paulnasdaq/fms-v2/common/users"
	"log"
	"time"
)

type AuthTokenService interface {
	Create(email string, password string) (string, error)
}

type authTokenImpl struct {
	repository      db.AuthRepository
	clients         *common.ClientManager
	userPassService *AuthUserPassService
	secretKey       []byte
}

func (a *authTokenImpl) Create(email string, password string) (string, error) {
	usersClient, err := a.clients.Users()
	if err != nil {
		log.Println("Error getting users service", err)
		return "", err
	}

	user, err := usersClient.GetByEmail(context.Context(context.Background()), &pb.GetByEmailRequest{Email: email})
	if err != nil {
		log.Println("Failed to get users", err)
		return "", err
	}
	_, err = a.userPassService.Verify(user.Id, password)
	if err != nil {
		log.Println("Failed to verify")
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.Id,
		"name": user.Email,
		"exp":  time.Now().Add(time.Hour * 24),
	})
	return token.SignedString(a.secretKey)
}

func NewAuthTokenService(repository db.AuthRepository, clients *common.ClientManager, userPassService *AuthUserPassService) AuthTokenService {
	return &authTokenImpl{repository: repository, secretKey: []byte("your-256-bit-secret"), clients: clients, userPassService: userPassService}
}
