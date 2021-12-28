package requests

import (
	"errors"
	"fmt"
)

type CreatePublisherReq struct {
	Name  string `json:"name"`
	Label string `json:"label"`
}

var maxPubNameLen = 50
var maxPubLabelLen = 50

func (req *CreatePublisherReq) Validate() error {
	if req.Name == "" {
		return errors.New("name is required")
	}
	if len(req.Name) > maxPubNameLen {
		return errors.New(fmt.Sprintf("name must be no more than %d characters", maxPubNameLen))
	}
	if req.Label == "" {
		return errors.New("label is required")
	}
	if len(req.Label) > maxPubLabelLen {
		return errors.New(fmt.Sprintf("label must be no more than %d characters", maxPubNameLen))
	}
	return nil
}
