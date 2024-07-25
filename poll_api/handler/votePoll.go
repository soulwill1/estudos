package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulwill1/estudos/models"
)

func VotePollOption(ctx *gin.Context) {
	pollID := ctx.Param("poll_id")
	optionID := ctx.Param("option_id")
	var option models.PollOption

	if err := db.Where("poll_id = ? AND id = ?", pollID, optionID).First(&option).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Option not found"})
		return
	}

	option.Votes++
	db.Save(&option)
	ctx.JSON(http.StatusOK, option)
}