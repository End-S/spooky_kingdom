package db

import (
	"fmt"

	"github.com/End-S/spooky_kingdom/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// New returns a new gorm database instance
func New() *gorm.DB {
	config := config.Get()

	db, err := gorm.Open(postgres.Open(config.DSN), &gorm.Config{})

	if err != nil {
		panic(fmt.Errorf("Could not connect to the database: %s", err))
	}

	return db
}
