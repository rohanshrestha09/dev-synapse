package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/dev-synapse/models"
	"github.com/rohanshrestha09/dev-synapse/service/database"
)

// Publish Project godoc
//
//	@Summary	Publish project
//	@Tags		Project
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Project ID"
//	@Success	201	{object}	database.Response
//	@Router		/project/{id}/publish [post]
//	@Security	Bearer
func Publish(ctx *gin.Context) {
	project := ctx.MustGet("project").(models.Project)

	data := map[string]any{"Published": true}

	response, err := database.Update(project, data)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusCreated, response("Project published"))

}

// Unpublish Project godoc
//
//	@Summary	Unpublish project
//	@Tags		Project
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Project ID"
//	@Success	201	{object}	database.Response
//	@Router		/project/{id}/publish [delete]
//	@Security	Bearer
func Unpublish(ctx *gin.Context) {
	project := ctx.MustGet("project").(models.Project)

	data := map[string]any{"Published": false}

	response, err := database.Update(project, data)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusCreated, response("Project unpublished"))

}
