package models

import (
	"strings"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	"github.com/End-S/spooky_kingdom/controllers/requests"
	"github.com/End-S/spooky_kingdom/utils"
	"github.com/google/uuid"
)

// ArticleModel struct
type ArticleModel struct {
	db *gorm.DB
}

// NewArticleModel creates a new article model instance
func NewArticleModel(db *gorm.DB) *ArticleModel {
	return &ArticleModel{
		db: db,
	}
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
	Accepted      bool
}

// List function returns a list of articles fitting the request
func (am ArticleModel) List(req *requests.GetArticlesReq, assessPending bool) ([]Article, int64, error) {
	var (
		articles []Article
		count    int64
	)

	qry := am.db.Table("articles").Where("accepted = ?", !assessPending)

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
func (am ArticleModel) Update(req *requests.UpdateArticleReq) (*Article, error) {
	var article Article

	res := am.db.Table("articles").Where("article_id = ?", req.ID).
		Updates(Article{Description: req.Description, ArticleType: req.Subject, Accepted: req.Accepted})

	if res.Error != nil {
		return nil, res.Error
	}

	am.db.Table("articles").First(&article, req.ID)

	return &article, res.Error
}

// Store stoes an article and returns the stored article
func (am ArticleModel) Store(req *requests.StoreArticleReq) (*Article, error) {
	var article Article
	var pub Publisher

	am.db.Table("publishers").Where("name = ?", req.PublisherName).First(&pub)

	if pub.Name == "" {
		// publisher not found create one
		uuid := uuid.New()
		am.db.Table("publishers").Create(Publisher{
			PublisherID: uuid,
			Name:        req.PublisherName,
			Label:       strings.Title(strings.ToLower(req.PublisherName)),
		})
		am.db.Table("publishers").First(&pub, "publisher_id = ?", uuid)
	}

	uuid := uuid.New()
	// store the article
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
	})

	if res.Error != nil {
		return nil, res.Error
	}

	am.db.Table("articles").Where("article_id = ?", uuid).First(&article)

	return &article, nil
}

// Delete removes the article from the database
func (am ArticleModel) Delete(id uuid.UUID) (int64, error) {
	res := am.db.Table("articles").Delete(&Article{}, id)

	return res.RowsAffected, res.Error
}
