package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulwill1/estudos/models"
)

func CreatePoll(ctx *gin.Context) {
	var poll models.Poll
	if err := ctx.ShouldBindJSON(&poll); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(poll.Options) < 3 || len(poll.Options) > 5 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "A poll must have between 3 and 5 options"})
		return
	}

	db.Create(&poll)
	ctx.JSON(http.StatusOK, poll)
}