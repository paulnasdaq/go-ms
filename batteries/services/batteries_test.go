package services

import (
	"github.com/paulnasdaq/fms-v2/batteries/db"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBatteryServiceImpl_Add(t *testing.T) {
	repo, _ := db.NewRepository()
	service, _ := NewBatteriesService(repo)
	serialNumber := "randomSerialNumber"
	t.Run("Test battery add", func(t *testing.T) {
		bat, err := service.Add(serialNumber)
		assert.NoError(t, err)
		assert.NotNil(t, bat)
	})
}
