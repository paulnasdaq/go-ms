package db

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Controller struct {
	gorm.Model
	ID           uuid.UUID
	SerialNumber string
	ModelID      uuid.UUID
}

type ControllerModel struct {
	gorm.Model
	ID          uuid.UUID
	Name        string
	Controllers []Controller `gorm:"foreignKey:ModelID"`
}

type VehicleController struct {
	gorm.Model
	ID           uuid.UUID
	VehicleID    uuid.UUID
	ControllerID uuid.UUID
}
