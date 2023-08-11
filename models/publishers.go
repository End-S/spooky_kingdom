package models

import (
	"errors"

	"github.com/google/uuid"
	"github.com/r-rayns/spooky_kingdom/controllers/requests"
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

func (pm *PublisherModel) Create(req *requests.CreatePublisherReq) (*Publisher, error) {
	var pub Publisher

	uuid := uuid.New()
	pm.db.Table("publishers").Create(Publisher{
		PublisherID: uuid,
		Name:        req.Name,
		Label:       req.Label,
	})
	pm.db.Table("publishers").First(&pub, "publisher_id = ?", uuid)

	if pub.Name == "" {
		return nil, errors.New("Could not create publisher")
	}

	return &pub, nil
}
