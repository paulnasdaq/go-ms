package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type UserPass struct {
	gorm.Model
	UserID   string `gorm:"primaryKey"`
	Password string
}

type UserPassRepository struct {
	db *gorm.DB
}

func NewUserPassRepository(db *gorm.DB) (*UserPassRepository, error) {
	if err := db.AutoMigrate(&UserPass{}); err != nil {
		log.Println("Failed to migrate UserPass table", err)
		return nil, err
	}
	return &UserPassRepository{db}, nil
}

func (r *UserPassRepository) Add(userID uuid.UUID, password string) error {
	if res := r.db.Create(&UserPass{UserID: userID.String(), Password: password}); res.Error != nil {
		log.Println("Failed to save password", res.Error)
		return res.Error
	}
	return nil
}

func (r *UserPassRepository) Get(userID uuid.UUID) (*UserPass, error) {
	var userPass UserPass
	if res := r.db.First(&userPass, &UserPass{UserID: userID.String()}); res.Error != nil {
		log.Println("Failed to get userPass", res.Error)
		return nil, res.Error
	}
	return &userPass, nil
}
