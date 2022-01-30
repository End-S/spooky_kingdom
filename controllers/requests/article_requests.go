package requests

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/End-S/spooky_kingdom/utils"
	"github.com/google/uuid"
)

// GetArticlesReq struct for a GetArticles request
type GetArticlesReq struct {
	Publishers    []uuid.UUID `query:"pbs"`
	FromDate      int64       `query:"frm"`
	ToDate        int64       `query:"to"`
	Order         string      `query:"ord"`
	SortBy        string      `query:"srt"`
	Subject       string      `query:"sbj"`
	Offset        int         `query:"oft"`
	Limit         int         `query:"lmt"`
	AssessPending string      `query:"pnd"`
}

// UpdateArticleReq struct for an UpdateArticles request
type UpdateArticleReq struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Subject     string    `json:"subject"`
	Accepted    bool      `json:"accepted"`
}

// StoreArticleReq struct for a StoreArticle request
type StoreArticleReq struct {
	PublisherName string `json:"publisherName"`
	DatePublished int64  `json:"datePublished"`
	DateRetrieved int64  `json:"dateRetrieved"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	Link          string `json:"link"`
	Subject       string `json:"subject"`
}

var validSubjects = []string{"ghost", "ufo", "abc"}
var maxDescLen = 500
var maxTitleLen = 2000
var maxLinkLen = 2000

func validSubject(sbj string) bool {
	return utils.Contains(validSubjects, sbj)
}

func (req *GetArticlesReq) Validate() error {
	if (req.FromDate != 0 && req.ToDate == 0) || (req.ToDate != 0 && req.FromDate == 0) {
		return errors.New("frm and to dates must be supplied together")
	}
	if (req.FromDate != 0 && req.ToDate != 0) && req.FromDate > req.ToDate {
		return errors.New("frm date must be before to date")
	}
	if req.Order == "" {
		return errors.New("ord is required")
	}
	if utils.Contains([]string{"asc", "desc"}, req.Order) == false {
		return errors.New("ord must be set to asc or desc")
	}
	if req.SortBy == "" {
		return errors.New("srt is required")
	}
	if utils.Contains([]string{"date", "title", "description"}, req.SortBy) == false {
		return errors.New("srt must be set to one of the following date, title or description")
	}
	if req.Subject != "" && validSubject(req.Subject) == false {
		return errors.New(fmt.Sprintf("Subject must be set to one of the following: %v", validSubjects))
	}
	if req.Limit == 0 {
		return errors.New("lmt is required")
	}
	if utils.Contains([]string{"10", "25", "50"}, strconv.Itoa(req.Limit)) == false {
		return errors.New("lmt must be set to one of the following 10, 25 or 50")
	}
	if req.Offset != 0 && req.Offset%req.Limit != 0 {
		return errors.New("oft must be a multiple of lmt")
	}
	if req.AssessPending != "" && utils.Contains([]string{"true", "false"}, req.AssessPending) == false {
		return errors.New("pnd must be set to one of the following true or false")
	}
	return nil
}

func (req *UpdateArticleReq) Validate() error {
	if req.Description == "" {
		return errors.New("description is required")
	}
	if len(req.Description) > maxDescLen {
		return errors.New(fmt.Sprintf("description must be less than %d characters", maxDescLen))
	}
	if req.Subject == "" {
		return errors.New("subject is required")
	}
	if req.Subject != "" && validSubject(req.Subject) == false {
		return errors.New(fmt.Sprintf("subject must be set to one of the following: %v", validSubjects))
	}
	return nil
}

func (req *StoreArticleReq) Validate() error {
	if req.PublisherName == "" {
		return errors.New("publisherName is required")
	}
	if len(req.PublisherName) > maxPubNameLen {
		return errors.New(fmt.Sprintf("publisherName must be no more than %d characters", maxPubNameLen))
	}
	if req.DatePublished == 0 {
		return errors.New("datePublished is required")
	}
	if req.DateRetrieved == 0 {
		return errors.New("dateRetrieved is required")
	}
	if req.Title == "" {
		return errors.New("title is required")
	}
	if len(req.Title) > maxTitleLen {
		return errors.New(fmt.Sprintf("title must be less than %d characters", maxTitleLen))
	}
	if req.Description == "" {
		return errors.New("description is required")
	}
	if len(req.Description) > maxDescLen {
		return errors.New(fmt.Sprintf("description must be less than %d characters", maxDescLen))
	}
	if req.Link == "" {
		return errors.New("link is required")
	}
	if len(req.Link) > maxLinkLen {
		return errors.New(fmt.Sprintf("link must be less than %d characters", maxLinkLen))
	}
	if req.Subject == "" {
		return errors.New("subject is required")
	}
	if req.Subject != "" && validSubject(req.Subject) == false {
		return errors.New(fmt.Sprintf("subject must be set to one of the following: %v", validSubjects))
	}
	return nil
}
