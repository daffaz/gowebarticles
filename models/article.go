package models

import "time"

type Article struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

var ArticleList []Article = []Article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body", CreatedAt: time.Now()},
	{ID: 2, Title: "Article 2", Content: "Article 2 body", CreatedAt: time.Now()},
}

func GetAllArticles() []Article {
	return ArticleList
}
