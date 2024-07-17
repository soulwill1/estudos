package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulwill1/estudos/schemas"
)

// @BasePath /api/v1

// @Summary show opening
// @Description show a job opening
// @Tags openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Sucess 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [get]

func ShowOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	if err := db.First(&opening, id).Error
	err != nil {
		sendError(ctx, http.StatusNotFound, "opening not found")
		return
	}

	sendSucess(ctx, "show-opening", opening)
}