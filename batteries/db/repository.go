package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

type Repository interface {
	Batteries() BatteryRepository
}

type mainRepositoryImpl struct {
	db                *gorm.DB
	batteryRepository BatteryRepository
}

func (m *mainRepositoryImpl) Batteries() BatteryRepository {
	return m.batteryRepository
}

func NewRepository() (Repository, error) {
	db, err := gorm.Open(sqlite.Open("batteries.db"), &gorm.Config{})
	if err != nil {
		log.Println("failed to create repo", err)
		return nil, err
	}
	batteryRepository, err := NewBatteryRepository(db)
	if err != nil {
		log.Println("failed to create repo", err)
		return nil, err
	}
	repo := mainRepositoryImpl{
		db:                db,
		batteryRepository: batteryRepository,
	}
	return &repo, nil
}
