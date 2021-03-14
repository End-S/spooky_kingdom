package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PublisherModel struct
type PublisherModel struct {
	db *gorm.DB
}

// NewPublisherModel creates a new publisher model instance
func NewPublisherModel(db *gorm.DB) *PublisherModel {
	return &PublisherModel{
		db: db,
	}
}

// Publisher struct
type Publisher struct {
	// gorm.Model
	PublisherID uuid.UUID
	Name        string
	Label       string
	Lat         float64
	Lng         float64
}

// List function returns a list of all publishers
func (pm *PublisherModel) List() ([]Publisher, error) {
	var publishers []Publisher

	res := pm.db.Find(&publishers)

	if res.Error != nil {
		return nil, res.Error
	}

	return publishers, nil
}
