package responses

// DeleteResponse json response for deletion requests
type DeleteResponse struct {
	Success bool `json:"success"`
}

// NewDeleteResponse creates a new DeleteResponse
func NewDeleteResponse(success bool) *DeleteResponse {
	res := new(DeleteResponse)
	res.Success = success

	return res
}
