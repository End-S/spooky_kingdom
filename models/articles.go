package models

import (
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/google/uuid"
	"github.com/r-rayns/spooky_kingdom/controllers/requests"
	"github.com/r-rayns/spooky_kingdom/utils"
)

// ArticleModel struct
type ArticleModel struct {
	db *gorm.DB
	pm *PublisherModel
}

// NewArticleModel creates a new article model instance
func NewArticleModel(db *gorm.DB, pm *PublisherModel) *ArticleModel {
	return &ArticleModel{db, pm}
}

// Article struct
type Article struct {
	// gorm.Model
	ArticleID     uuid.UUID
	PublisherID   uuid.UUID
	Publisher     string
	DatePublished time.Time
	DateRetrieved time.Time
	Title         string
	Description   string
	Link          string
	ArticleType   string
	ArticleState  string
	MatchedTerms  string
}

type DateSpan struct {
	Max time.Time
	Min time.Time
}

// List function returns a list of articles fitting the request
func (am *ArticleModel) List(req *requests.GetArticlesReq, assessPending bool) ([]Article, int64, error) {
	var (
		articles []Article
		count    int64
	)

	articleState := "accepted"

	if assessPending {
		// get articles that are pending review
		articleState = "review"
	}

	qry := am.db.Table("articles").Where("article_state= ?", articleState)

	if req.Subject != "" {
		// filter by subject
		qry.Where("article_type = ?", req.Subject)
	}

	if len(req.Publishers) > 0 {
		// filter by publisher id
		qry.Where("publisher_id IN ?", req.Publishers)
	}

	if req.FromDate > 0 && req.ToDate > req.FromDate {
		// filter by date range
		from := time.Unix(req.FromDate, 0)
		to := time.Unix(req.ToDate, 0)
		qry.Where("date_published BETWEEN ? AND ?", from, to)
	}

	// order results by column
	validSorts := [3]string{"title", "description", "date_publised"}
	if utils.Contains(validSorts[0:], req.SortBy) {
		qry.Order(clause.OrderByColumn{Column: clause.Column{Name: req.SortBy}, Desc: req.Order == "desc"})
	} else {
		qry.Order(clause.OrderByColumn{Column: clause.Column{Name: "date_published"}, Desc: req.Order == "desc"})
	}

	// count affected
	qry.Model(&articles).Count(&count)
	// paginate results
	qry.Offset(req.Offset).
		Limit(req.Limit).Find(&articles)

	return articles, count, nil
}

// Update updates an article and returns the updated article
func (am *ArticleModel) Update(id uuid.UUID, req *requests.UpdateArticleReq) (*Article, error) {
	var article Article

	articleState := "accepted"

	if !req.Accepted {
		articleState = "rejected"
	}

	res := am.db.Table("articles").Where("article_id = ?", id).
		Updates(Article{Description: req.Description, ArticleType: req.Subject, ArticleState: articleState})

	if res.Error != nil {
		return nil, res.Error
	}

	am.db.Table("articles").First(&article, id)

	return &article, res.Error
}

// Store stores an article and returns the stored article
func (am *ArticleModel) Store(req *requests.StoreArticleReq) (*Article, error) {
	var article Article
	var pub Publisher

	am.db.Table("publishers").Where("name = ?", req.PublisherName).First(&pub)

	if pub.Name == "" {
		// publisher not found create one
		newPub, err := am.pm.Create(&requests.CreatePublisherReq{
			Name:  req.PublisherName,
			Label: strings.Title(strings.ToLower(req.PublisherName)),
		})
		if err != nil {
			return nil, err
		}
		pub = *newPub
	}

	uuid := uuid.New()
	// store the article, will fail if duplicate title is found
	res := am.db.Table("articles").Create(Article{
		ArticleID:     uuid,
		PublisherID:   pub.PublisherID,
		Publisher:     pub.Name,
		DatePublished: time.Unix(req.DatePublished, 0),
		DateRetrieved: time.Unix(req.DateRetrieved, 0),
		Title:         req.Title,
		Description:   req.Description,
		Link:          req.Link,
		ArticleType:   req.Subject,
		ArticleState:  "review",
		MatchedTerms:  strings.Join(req.MatchedTerms, ","),
	})

	if res.Error != nil {
		return nil, res.Error
	}

	am.db.Table("articles").Where("article_id = ?", uuid).First(&article)

	return &article, nil
}

// Delete removes the article from the database
func (am *ArticleModel) Delete(id uuid.UUID) (int64, error) {
	res := am.db.Table("articles").Delete(&Article{}, id)

	return res.RowsAffected, res.Error
}

// DateSpan gets the max (latest) and min (oldest) article dates
func (am *ArticleModel) DateSpan() (*DateSpan, error) {
	var dateSpan DateSpan
	res := am.db.Raw("SELECT max(date_published), min(date_published) FROM articles WHERE article_state = 'accepted'").Scan(&dateSpan)

	if res.Error != nil {
		return nil, res.Error
	}

	return &dateSpan, nil
}
