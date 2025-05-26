package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type Battery struct {
	gorm.Model
	ID           uuid.UUID
	SerialNumber string
}

type BatteryRepository interface {
	Add(serialNumber string) (*Battery, error)
	Get(ID uuid.UUID) (*Battery, error)
	init() error
}

type BatteryRepositoryImpl struct {
	db *gorm.DB
}

func (b *BatteryRepositoryImpl) init() error {
	return b.db.AutoMigrate(&Battery{})
}

func (b *BatteryRepositoryImpl) Add(serialNumber string) (*Battery, error) {
	newBattery := Battery{
		SerialNumber: serialNumber,
		ID:           uuid.New(),
	}
	if res := b.db.Create(&newBattery); res.Error != nil {
		log.Println("Error creating battery: ", res.Error)
		return nil, res.Error
	}
	return &newBattery, nil
}

func (b *BatteryRepositoryImpl) Get(ID uuid.UUID) (*Battery, error) {
	var battery Battery
	if res := b.db.First(&battery, ID); res.Error != nil {
		log.Println("Failed to get battery: ", res.Error)
		return nil, res.Error
	}
	return &battery, nil
}

func NewBatteryRepository(db *gorm.DB) (BatteryRepository, error) {
	if err := db.AutoMigrate(&Battery{}); err != nil {
		log.Println("Failed to initialize battery repo", err)
		return nil, err
	}
	return &BatteryRepositoryImpl{db: db}, nil
}
