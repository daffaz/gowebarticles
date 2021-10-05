package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// Set the router using the default one provided by Gin
	router := gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*/*.html")

	// Route handler
	router.GET("/", func(ctx *gin.Context) {
		// Call the HTML method of the context to render a template
		ctx.HTML(http.StatusOK, "index.html", gin.H{"title": "Home Page"})
	})

	router.Run()
}
