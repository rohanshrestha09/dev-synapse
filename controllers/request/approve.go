package request

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/dev-synapse/enums"
	"github.com/rohanshrestha09/dev-synapse/models"
	"github.com/rohanshrestha09/dev-synapse/service/database"
)

// Approve Request godoc
//
//	@Summary	Approve request
//	@Tags		Request
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Request ID"
//	@Success	201	{object}	database.Response
//	@Router		/request/{id}/approve [post]
//	@Security	Bearer
func Approve(ctx *gin.Context) {
	request := ctx.MustGet("request").(models.Request)

	_, err := database.Create(models.Developer{
		UserID:    request.UserID,
		ProjectID: request.ProjectID,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	data := models.Request{Status: enums.APPROVED}

	response, err := database.Update(request, data)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusCreated, response("Request approved"))
}

// Reject Request godoc
//
//	@Summary	Reject request
//	@Tags		Request
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Request ID"
//	@Success	201	{object}	database.Response
//	@Router		/request/{id}/approve [delete]
//	@Security	Bearer
func Reject(ctx *gin.Context) {
	request := ctx.MustGet("request").(models.Request)

	data := models.Request{Status: enums.REJECTED}

	response, err := database.Update(request, data)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusCreated, response("Request rejected"))
}
