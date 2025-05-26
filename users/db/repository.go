package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Repository interface {
	Users() *UsersRepository
}

type repoImpl struct {
	usersRepo *UsersRepository
}

func (r repoImpl) Users() *UsersRepository {
	return r.usersRepo
}

func NewRepository() (Repository, error) {
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	usersRepo, err := NewUsersRepository(db)
	if err != nil {
		return nil, err
	}
	return &repoImpl{usersRepo: usersRepo}, nil
}
