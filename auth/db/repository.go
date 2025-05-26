package db

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type AuthRepository interface {
	UserPass() *UserPassRepository
	Token() *TokenRepository
}

type authRepositoryImpl struct {
	userPassRepository *UserPassRepository
	token              *TokenRepository
	db                 *gorm.DB
}

func (r *authRepositoryImpl) Token() *TokenRepository {
	return r.token
}

func (r *authRepositoryImpl) UserPass() *UserPassRepository {
	return r.userPassRepository
}

func NewRepository() (AuthRepository, error) {
	db, err := gorm.Open(sqlite.Open("auth.db"), &gorm.Config{})
	if err != nil {
		log.Println("Failed to create db", err)
		return nil, err
	}
	userPassRepository, err := NewUserPassRepository(db)
	if err != nil {
		return nil, err
	}
	cache := redis.NewClient(&redis.Options{Addr: "localhost:6379", Password: "", DB: 0})

	tokenRepository := NewTokenRepository(cache)

	return &authRepositoryImpl{db: db, userPassRepository: userPassRepository, token: tokenRepository}, nil
}
