package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitAndRunRoutes(router *gin.Engine) {
	// Route handler
	router.GET("/", ShowIndexPage)
	router.Run(":8080")
}

func ShowIndexPage(ctx *gin.Context) {
	// Call the HTML method of the context to render a template
	ctx.HTML(http.StatusOK, "index.html", gin.H{"title": "Home Page"})
}
