package responses

import (
	"time"

	"github.com/End-S/spooky_kingdom/models"
	"github.com/google/uuid"
)

type articleResponse struct {
	ArticleID     uuid.UUID `json:"id"`
	Publisher     publisher `json:"publisher"`
	DatePublished time.Time `json:"datePublished"`
	DateRetrieved time.Time `json:"dateRetrieved"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	Link          string    `json:"link"`
	ArticleType   string    `json:"type"`
	Accepted      bool      `json:"accepted"`
}

type publisher struct {
	PublisherID   uuid.UUID `json:"id"`
	PublisherName string    `json:"name"`
}

// ListArticlesResponse json reponse containing a list of articles
type ListArticlesResponse struct {
	Articles []*articleResponse `json:"articles"`
	Count    int64              `json:"total"`
}

// NewListArticlesResponse creates a new ListArticlesResponse
func NewListArticlesResponse(articles []models.Article, count int64) *ListArticlesResponse {
	res := new(ListArticlesResponse)
	res.Articles = make([]*articleResponse, 0)
	for _, a := range articles {
		aRes := new(articleResponse)
		aRes.ArticleID = a.ArticleID
		aRes.Publisher.PublisherID = a.PublisherID
		aRes.Publisher.PublisherName = a.Publisher
		aRes.DatePublished = a.DatePublished
		aRes.DateRetrieved = a.DateRetrieved
		aRes.Title = a.Title
		aRes.Description = a.Description
		aRes.Link = a.Link
		aRes.ArticleType = a.ArticleType
		aRes.Accepted = a.Accepted
		res.Articles = append(res.Articles, aRes)
	}
	res.Count = count

	return res
}

// UpdateArticleResponse json reponse containing the updated article
type UpdateArticleResponse struct {
	Article *articleResponse `json:"article"`
}

// NewSingleArticleResponse creates a new UpdateArticleResponse
func NewSingleArticleResponse(article models.Article) *UpdateArticleResponse {
	res := new(UpdateArticleResponse)
	// res.Article = new(articleResponse)
	res.Article = &articleResponse{
		ArticleID: article.ArticleID,
		Publisher: publisher{
			PublisherID:   article.PublisherID,
			PublisherName: article.Publisher,
		},
		DatePublished: article.DatePublished,
		DateRetrieved: article.DateRetrieved,
		Title:         article.Title,
		Description:   article.Description,
		Link:          article.Link,
		ArticleType:   article.ArticleType,
		Accepted:      article.Accepted,
	}
	return res
}
