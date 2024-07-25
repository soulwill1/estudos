package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulwill1/estudos/models"
)

func DeletePoll(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := db.Delete(&models.Poll{}, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Poll not found"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Poll deleted"})
}