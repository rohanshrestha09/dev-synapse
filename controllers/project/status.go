package project

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/dev-synapse/enums"
	"github.com/rohanshrestha09/dev-synapse/models"
	"github.com/rohanshrestha09/dev-synapse/service/database"
)

// Update Project Status godoc
//
//	@Summary	Update project status
//	@Tags		Project
//	@Accept		json
//	@Produce	json
//	@Param		id		path		int		true	"Project ID"
//	@Param		status	query		string	true	"Status"	Enums(OPEN, IN_PROGRESS, CLOSED)
//	@Success	201		{object}	database.Response
//	@Router		/project/{id}/status [post]
//	@Security	Bearer
func UpdateStatus(ctx *gin.Context) {
	project := ctx.MustGet("project").(models.Project)

	var query struct {
		Status enums.ProjectStatus `form:"status"`
	}

	if err := ctx.BindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	data := models.Project{
		Status: query.Status,
	}

	switch query.Status {
	case enums.IN_PROGRESS:
		data.StartDate.Scan(time.Now())
		data.EndDate.Scan(time.Now().AddDate(0, 0, int(project.EstimatedDuration)))

	case enums.OPEN:
		data.StartDate.Valid = false
		data.EndDate.Valid = false
	}

	response, err := database.Update(project, data)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusCreated, response("Project status updated"))

}
