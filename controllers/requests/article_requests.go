package requests

import "github.com/google/uuid"

// GetArticlesReq struct for a GetArticles request
type GetArticlesReq struct {
	Publishers    []uuid.UUID `query:"pbs"`
	FromDate      int64       `query:"frm" validate:"required_with=ToDate,omitempty,ltfield=ToDate,numeric" ltfield:"From date must be before to date"`
	ToDate        int64       `query:"to" validate:"required_with=FromDate,omitempty,gtfield=FromDate,numeric" gtfield:"To date must be after from date"`
	Order         string      `query:"ord" validate:"required,oneof=asc desc" oneof:"Order must be set to asc or desc"`
	SortBy        string      `query:"srt" validate:"required,oneof=date title description" oneof:"Sort by must be set to date, title or description"`
	Subject       string      `query:"sbj" validate:"omitempty,valid_subject" valid_subject:"Subject must be set to a recognised subject"`
	Offset        int         `query:"oft" validate:"multiple_of=Limit" multiple_of:"Offset must be a multiple of Limit"`
	Limit         int         `query:"lmt" validate:"required,oneof=10 25 50" oneof:"Limit must be set to 10 25 or 50"`
	AssessPending string      `query:"pnd" validate:"omitempty,oneof=true false" oneof:"Assess pending must be set to true or false"`
}

// UpdateArticleReq struct for an UpdateArticles request
type UpdateArticleReq struct {
	ID          uuid.UUID `json:"id" validate:"required"`
	Description string    `json:"description" validate:"required,max=500" max:"Description cannot be over 500 character"`
	Subject     string    `json:"subject" validate:"required,valid_subject" valid_subject:"Subject must be set to a recognised subject"`
	Accepted    bool      `json:"accepted" validate:"required"`
}
