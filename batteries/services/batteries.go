package services

import (
	"github.com/paulnasdaq/fms-v2/batteries/db"
	"log"
	"time"
)

type Battery struct {
	ID           string
	SerialNumber string
	CreatedAt    time.Time
}
type BatteriesService interface {
	Add(serialNumber string) (*Battery, error)
	Get(ID string) (*Battery, error)
}

type batteryServiceImpl struct {
	repository db.Repository
}

func (b *batteryServiceImpl) Add(serialNumber string) (*Battery, error) {
	res, err := b.repository.Batteries().Add(serialNumber)
	if err != nil {
		log.Println("failed to add battery", err)
		return nil, err
	}
	return &Battery{ID: res.ID.String(), SerialNumber: res.SerialNumber, CreatedAt: res.CreatedAt}, nil
}

func (b *batteryServiceImpl) Get(ID string) (*Battery, error) {
	//TODO implement me
	panic("implement me")
}

func NewBatteriesService(repository db.Repository) (BatteriesService, error) {
	return &batteryServiceImpl{repository: repository}, nil
}
