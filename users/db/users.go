package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	ID    uuid.UUID
	Email string
}

type UsersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) (*UsersRepository, error) {
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Println("Failed to migrate users table", err)
		return nil, err
	}
	return &UsersRepository{db: db}, nil
}

func (r *UsersRepository) Add(email string) (*User, error) {
	user := User{Email: email, ID: uuid.New()}
	res := r.db.Create(&user)
	if res.Error != nil {
		log.Println("Failed to add user", res.Error)
		return nil, res.Error
	}
	return &user, nil
}

func (r *UsersRepository) Get(id uuid.UUID) (*User, error) {
	var user User
	if res := r.db.First(&user, id); res.Error != nil {
		log.Println("Failed to get user", res.Error)
		return nil, res.Error
	}
	return &user, nil
}
func (r *UsersRepository) GetByEmail(email string) (*User, error) {
	var user User
	res := r.db.First(&user, User{Email: email})
	if res.Error != nil {
		log.Println("Failed to find user ", res.Error)
		return nil, res.Error
	}
	return &user, nil

}
