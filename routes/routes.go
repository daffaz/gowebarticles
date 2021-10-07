package routes

import (
	"errors"
	"gowebmicro/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitAndRunRoutes(router *gin.Engine) {
	// Route handler
	router.GET("/", ShowIndexPage)
	router.GET("/article/view/:slug", GetArticle)
	router.Run(":8080")
}

func ShowIndexPage(ctx *gin.Context) {
	// Call the HTML method of the context to render a template
	ctx.HTML(http.StatusOK, "index.html", gin.H{"title": "Home Page", "payload": models.ArticleList})
}

func GetArticle(ctx *gin.Context) {
	article, err := FetchSingleArticle(ctx.Param("slug"))
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
	}
	ctx.HTML(http.StatusOK, "article.html", gin.H{"title": article.Title, "payload": article})
}

func FetchSingleArticle(slug string) (*models.Article, error) {
	for _, article := range models.ArticleList {
		if article.Slug == slug {
			return &article, nil
		}
	}
	return nil, errors.New("article not found")
}
