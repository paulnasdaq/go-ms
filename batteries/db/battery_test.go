package db

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

func setup() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db.AutoMigrate(&Battery{})

	return db
}

func TestBatteryRepositoryImpl_Add(t *testing.T) {
	db := setup()

	serialNumber := "random"
	t.Run("creating batter", func(t *testing.T) {
		testBatteriesRepo, _ := NewBatteryRepository(db)
		newBattery, err := testBatteriesRepo.Add(serialNumber)

		assert.NoError(t, err)
		assert.NotNil(t, newBattery)
		assert.NotNil(t, newBattery.ID)
	})

}

func TestBatteryRepositoryImpl_Get(t *testing.T) {
	db := setup()
	serialNumber := "random"
	t.Run("Test get", func(t *testing.T) {
		testBatteriesRepo, _ := NewBatteryRepository(db)
		newBatt, _ := testBatteriesRepo.Add(serialNumber)

		bat, err := testBatteriesRepo.Get(newBatt.ID)

		assert.NotNil(t, bat)
		assert.NoError(t, err)

		bat, err = testBatteriesRepo.Get(uuid.New())

		assert.Error(t, err)
		assert.Nil(t, bat)
	})
}
