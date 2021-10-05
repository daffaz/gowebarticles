package main

import (
	"gowebmicro/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set the router using the default one provided by Gin
	router := gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*/*.html")

	// Initialize the routes + running it
	routes.InitAndRunRoutes(router)
}
