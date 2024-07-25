package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/soulwill1/estudos/models"
)

func GetPolls(ctx *gin.Context) {
	// Esta variável armazena as enquetes (polls) que serão recuperadas do banco de dados
	var polls []models.Poll

	// Preload("Options") carrega as opções associadas a cada enquete
	// Find(&polls) recupera todas as enquetes do banco de dados e armazena na variável polls
	db.Preload("Options").Find(&polls)
	now := time.Now()

	for i := range polls {
		if now.After(polls[i].StartDate) && now.Before(polls[i].EndDate) {
			polls[i].Status = "active"
		} else if now.Before(polls[i].StartDate) {
			polls[i].Status = "not_started"
		} else {
			polls[i].Status = "ended"
		}
	}

	// Retorna as enquetes e seus status em formato JSON com código HTTP 200 (OK)
	ctx.JSON(http.StatusOK, polls)
}
