package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/dev-synapse/models"
	"github.com/rohanshrestha09/dev-synapse/service/database"
)

// Get Developers godoc
//
//	@Summary	Get project developers
//	@Tags		Project
//	@Accept		json
//	@Produce	json
//	@Param		id		path		int		true	"Project ID"
//	@Param		sort	query		string	false	"Sort"	Enums(id, created_at)
//	@Param		order	query		string	false	"Order"	Enums(asc, desc)
//	@Success	200		{object}	database.GetAllResponse[models.Developer]
//	@Router		/project/{id}/developer [get]
func GetDevelopers(ctx *gin.Context) {

	project := ctx.MustGet("project").(models.Project)

	args := database.GetAllArgs[models.Developer]{
		Include: map[string][]string{
			"User": {"Email"},
		},
		Filter: models.Developer{
			ProjectID: project.ID,
		},
	}

	response, err := database.GetAll(ctx.BindQuery, args)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}
