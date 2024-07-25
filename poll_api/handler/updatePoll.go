package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulwill1/estudos/models"
)

func UpdatePoll(ctx *gin.Context) {
	id := ctx.Param("id")
	var poll models.Poll
	if err := db.First(&poll, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Poll not found"})
		return
	}

	if err := ctx.ShouldBindJSON(&poll); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&poll)
	ctx.JSON(http.StatusOK, poll)
}