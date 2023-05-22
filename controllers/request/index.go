package request

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/dev-synapse/enums"
	"github.com/rohanshrestha09/dev-synapse/models"
	"github.com/rohanshrestha09/dev-synapse/service/database"
)

// Get Request godoc
//
//	@Summary	Get a request
//	@Tags		Request
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Request ID"
//	@Success	200	{object}	database.GetResponse[models.Request]
//	@Router		/request/{id} [get]
func Get(ctx *gin.Context) {
	request := ctx.MustGet("request").(models.Request)

	ctx.JSON(
		http.StatusOK,
		database.GetResponse[models.Request]{
			Message: "Request Fetched",
			Data:    request,
		})
}

// Create Request godoc
//
//	@Summary	Create request
//	@Tags		Request
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Project ID"
//	@Success	201	{object}	database.Response
//	@Router		/request/{id} [post]
//	@Security	Bearer
func Create(ctx *gin.Context) {
	project := ctx.MustGet("project").(models.Project)

	authUser := ctx.MustGet("auth").(models.User)

	if recordExists, err := database.RecordExists(models.Request{
		ProjectID: project.ID,
		UserID:    authUser.ID,
	}); recordExists && err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	} else if recordExists {
		ctx.JSON(http.StatusForbidden, gin.H{"message": "Request already exists"})
		return
	}

	request := models.Request{
		ProjectID: project.ID,
		UserID:    authUser.ID,
	}

	response, err := database.Create(request)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusCreated, response("Request Created"))
}

// Delete Request godoc
//
//	@Summary	Delete request
//	@Tags		Request
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Request ID"
//	@Success	201	{object}	database.Response
//	@Router		/request/{id} [delete]
//	@Security	Bearer
func Delete(ctx *gin.Context) {
	request := ctx.MustGet("request").(models.Request)

	if request.Status == enums.APPROVED {
		ctx.JSON(http.StatusBadRequest, database.Response{Message: "cannot delete approved request"})
		return
	}

	response, err := database.Delete(request)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusCreated, response("Request Deleted"))
}
