package responses

import (
	"github.com/End-S/spooky_kingdom/models"
	"github.com/google/uuid"
)

type publisherResponse struct {
	PublisherID uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Label       string    `json:"label"`
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
		res.Publishers = append(res.Publishers, pRes)
	}
	return res
}

type PublisherResponseBody struct {
	PublisherID   uuid.UUID `json:"id"`
	PublisherName string    `json:"name"`
}

type PublisherResponse struct {
	Publisher *PublisherResponseBody `json:"publisher"`
}

func NewPublisherResponse(publisher models.Publisher) *PublisherResponse {
	res := new(PublisherResponse)
	res.Publisher = &PublisherResponseBody{
		PublisherID:   publisher.PublisherID,
		PublisherName: publisher.Name,
	}

	return res
}
