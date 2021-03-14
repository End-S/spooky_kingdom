package responses

import (
	"github.com/End-S/spooky_kingdom/models"
	"github.com/google/uuid"
)

type publisherResponse struct {
	PublisherID uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Label       string     `json:"label"`
	LatLng      [2]float64 `json:"latlng"`
}

// ListPublishersResponse json response containing a list of publishers
type ListPublishersResponse struct {
	Publishers []*publisherResponse `json:"publishers"`
}

// NewListPublishersResponse creates a new ListPublishersResponse
func NewListPublishersResponse(publishers []models.Publisher) *ListPublishersResponse {
	res := new(ListPublishersResponse)
	res.Publishers = make([]*publisherResponse, 0)
	for _, p := range publishers {
		pRes := new(publisherResponse)
		pRes.PublisherID = p.PublisherID
		pRes.Name = p.Name
		pRes.Label = p.Label
		pRes.LatLng = [2]float64{p.Lat, p.Lng}
		res.Publishers = append(res.Publishers, pRes)
	}
	return res
}
