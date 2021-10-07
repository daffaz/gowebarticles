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

func Render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data)
	case "application/xml":
		c.XML(http.StatusOK, data)
	default:
		c.HTML(http.StatusOK, templateName, data)
	}
}

func ShowIndexPage(ctx *gin.Context) {
	// Call the HTML method of the context to render a template
	Render(ctx, gin.H{"title": "Home Page", "payload": models.ArticleList}, "index.html")
}

func GetArticle(ctx *gin.Context) {
	article, err := FetchSingleArticle(ctx.Param("slug"))
	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
	}
	Render(ctx, gin.H{
		"title":   article.Title,
		"payload": article,
	}, "article.html")
}

func FetchSingleArticle(slug string) (*models.Article, error) {
	for _, article := range models.ArticleList {
		if article.Slug == slug {
			return &article, nil
		}
	}
	return nil, errors.New("article not found")
}
