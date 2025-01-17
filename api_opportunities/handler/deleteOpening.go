package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/soulwill1/estudos/schemas"
)

// @BasePath /api/v1

// @Summary delete opening
// @Description delete a job opening
// @Tags openings
// @Accept json
// @Produce json
// @Param id query string true "Opening identification"
// @Sucess 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /opening [delete]

func DeleteOpeningHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	opening := schemas.Opening{}
	// Find opening
	if err := db.First(&opening, id).Error
	err != nil {
		sendError(ctx, http.StatusNotFound, fmt.Sprintf("opening with id: %s not found", id))
		return
	}
	// Delete opening
	if err := db.Delete(&opening).Error
	err != nil {
		sendError(ctx, http.StatusInternalServerError, fmt.Sprintf("error deleting opening with id: %s", id))
		return
	}

	sendSucess(ctx, "delete-opening", opening)
}