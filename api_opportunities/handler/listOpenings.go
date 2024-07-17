package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulwill1/estudos/schemas"
)

// @BasePath /api/v1

// @Summary list opening
// @Description list a new job opening
// @Tags openings
// @Accept json
// @Produce json
// @Sucess 200 {object} ListOpeningsResponse
// @Failure 500 {object} ErrorResponse
// @Router /opening [get]

func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error

	err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}

	sendSucess(ctx, "list-openings", openings)
}