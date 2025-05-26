package services

import (
	"github.com/google/uuid"
	"github.com/paulnasdaq/fms-v2/auth/db"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type AuthUserPassService struct {
	repository db.AuthRepository
}

func NewAuthUserPassService(repository db.AuthRepository) *AuthUserPassService {
	return &AuthUserPassService{repository: repository}
}
func (s *AuthUserPassService) Add(userID string, password string) error {
	userIDUUID, err := uuid.Parse(userID)
	if err != nil {

		return err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Failed to hash password", err)
		return err
	}
	err = s.repository.UserPass().Add(userIDUUID, string(hash))
	if err != nil {
		log.Println("Failed to add password", err)
		return err
	}
	return nil
}

func (s *AuthUserPassService) Verify(userID string, password string) (bool, error) {
	userIDUUID, err := uuid.Parse(userID)
	if err != nil {
		return false, err
	}
	hash, err := s.repository.UserPass().Get(userIDUUID)
	if err != nil {
		log.Println("Failed to get userpass", err)
		return false, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hash.Password), []byte(password))
	if err != nil {
		return false, err
	}
	return true, nil
}
