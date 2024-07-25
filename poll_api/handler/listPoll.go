package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulwill1/estudos/models"
)

func GetPollByID(ctx *gin.Context) {
	id := ctx.Param("id")
	var poll models.Poll
	if err := db.Preload("Options").First(&poll, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Poll not found"})
		return
	}
	ctx.JSON(http.StatusOK, poll)
}