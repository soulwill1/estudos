package router

import (
	"github.com/gin-gonic/gin"
	"github.com/soulwill1/estudos/handler"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/swag/example/basic/docs"
)
// @BasePath /api/v1

// @Summary create opening
// @Description create a new job opening
// @Tags openings
// @Accept json
// @Produce json
// @Param request body CreateOpeningRequest true "Request body"
// @Sucess 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [post]
 

func initializeRoutes(router *gin.Engine) {
	// Initiliaze handler
	handler.InitializeHandler()
	basePath := "api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group("/api/v1")

	{
		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.GET("/openings", handler.ListOpeningsHandler)
	}
	// Initialize swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}