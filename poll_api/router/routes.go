package router

import (
	"github.com/gin-gonic/gin"
	"github.com/soulwill1/estudos/handler"
)

func RoutesInit(router *gin.Engine) {
	handler.InitializeHandler()
	
	v1 := router.Group("/api/v1")

	{
		v1.POST("/polls", handler.CreatePoll)
		v1.GET("/polls", handler.GetPolls)
		v1.GET("/polls/:id", handler.GetPollByID)
		v1.PUT("/polls/:id", handler.UpdatePoll)
		v1.DELETE("/polls/:id", handler.DeletePoll)
		v1.POST("/polls/:poll_id/options/:option_id/vote", handler.VotePollOption)
	}

}