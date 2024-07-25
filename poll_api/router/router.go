package router

import "github.com/gin-gonic/gin"

func RouterInit() {
	router := gin.Default()

	RoutesInit(router)

	router.Run(":8080")
}