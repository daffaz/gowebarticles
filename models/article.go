package models

import (
	"time"

	"github.com/gosimple/slug"
)

type Article struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Slug      string    `json:"slug"`
	CreatedAt time.Time `json:"created_at"`
}

var ArticleList []Article = []Article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body", Slug: slug.Make("Article 1"), CreatedAt: time.Now()},
	{ID: 2, Title: "Article 2", Content: "Article 2 body", Slug: slug.Make("Article 2"), CreatedAt: time.Now()},
}

func GetAllArticles() []Article {
	return ArticleList
}
