package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize Router
	router := gin.Default()

	// Initialize Routes
	initializeRoutes(router)

	// Initialize Server
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}