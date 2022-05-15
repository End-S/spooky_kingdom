package responses

import (
	"strings"
	"time"

	"github.com/End-S/spooky_kingdom/models"
	"github.com/google/uuid"
)

type ArticleResponseBody struct {
	ArticleID     uuid.UUID             `json:"id"`
	Publisher     PublisherResponseBody `json:"publisher"`
	DatePublished time.Time             `json:"datePublished"`
	DateRetrieved time.Time             `json:"dateRetrieved"`
	Title         string                `json:"title"`
	Description   string                `json:"description"`
	Link          string                `json:"link"`
	ArticleType   string                `json:"type"`
	ArticleState  string                `json:"state"`
	MatchedTerms  []string              `json:"matchedTerms"`
}

type ListArticlesResponse struct {
	Articles []*ArticleResponseBody `json:"articles"`
	Count    int64                  `json:"total"`
}

type ArticleDateSpanResponse struct {
	MaxDate time.Time `json:"max"`
	MinDate time.Time `json:"min"`
}

// NewListArticlesResponse creates a new ListArticlesResponse
func NewListArticlesResponse(articles []models.Article, count int64) *ListArticlesResponse {
	res := new(ListArticlesResponse)
	res.Articles = make([]*ArticleResponseBody, 0)
	for _, a := range articles {
		aRes := new(ArticleResponseBody)
		aRes.ArticleID = a.ArticleID
		aRes.Publisher.PublisherID = a.PublisherID
		aRes.Publisher.PublisherName = a.Publisher
		aRes.DatePublished = a.DatePublished
		aRes.DateRetrieved = a.DateRetrieved
		aRes.Title = a.Title
		aRes.Description = a.Description
		aRes.Link = a.Link
		aRes.ArticleType = a.ArticleType
		aRes.ArticleState = a.ArticleState
		aRes.MatchedTerms = strings.Split(a.MatchedTerms, ",")
		res.Articles = append(res.Articles, aRes)
	}
	res.Count = count

	return res
}

// ArticleResponse json response containing the article
type ArticleResponse struct {
	Article *ArticleResponseBody `json:"article"`
}

// NewArticleResponse creates a new UpdateArticleResponse
func NewArticleResponse(article models.Article) *ArticleResponse {
	res := new(ArticleResponse)
	res.Article = &ArticleResponseBody{
		ArticleID: article.ArticleID,
		Publisher: PublisherResponseBody{
			PublisherID:   article.PublisherID,
			PublisherName: article.Publisher,
		},
		DatePublished: article.DatePublished,
		DateRetrieved: article.DateRetrieved,
		Title:         article.Title,
		Description:   article.Description,
		Link:          article.Link,
		ArticleType:   article.ArticleType,
		ArticleState:  article.ArticleState,
		MatchedTerms:  strings.Split(article.MatchedTerms, ","),
	}
	return res
}

// NewArticleDateSpanResponse creates a new ArticleDateSpanResponse
func NewArticleDateSpanResponse(dateRange models.DateSpan) *ArticleDateSpanResponse {
	res := new(ArticleDateSpanResponse)
	res.MaxDate = dateRange.Max
	res.MinDate = dateRange.Min

	return res
}
