package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
)

type Vehicle struct {
	gorm.Model
	ID            uuid.UUID
	RegNumber     string
	ChassisNumber string
	ModelID       uuid.UUID
}

type VehicleUser struct {
	gorm.Model
	ID        uuid.UUID
	VehicleID uuid.UUID
	UserID    uuid.UUID
}

type VehiclesDomain struct {
	db *gorm.DB
}

func (d *VehiclesDomain) Add(regNumber string, chassisNumber string, modelID uuid.UUID) (*Vehicle, error) {
	vehicle := Vehicle{ID: uuid.New(), RegNumber: regNumber, ChassisNumber: chassisNumber, ModelID: modelID}
	res := d.db.Create(&vehicle)
	if res.Error != nil {
		log.Println(res.Error)
		return nil, res.Error
	}
	return &vehicle, nil
}

func (d *VehiclesDomain) Get(ID uuid.UUID) (*Vehicle, error) {
	vehicle := Vehicle{}
	if res := d.db.First(&vehicle, ID); res.Error != nil {
		return nil, res.Error
	}
	return &vehicle, nil
}

func (d *VehiclesDomain) Delete(ID uuid.UUID) error {
	if res := d.db.Delete(&Vehicle{}, ID); res.Error != nil {
		return res.Error
	}
	return nil
}
