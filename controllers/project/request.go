package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/dev-synapse/enums"
	"github.com/rohanshrestha09/dev-synapse/models"
	"github.com/rohanshrestha09/dev-synapse/service/database"
)

// Get Requests godoc
//
//	@Summary	Get project requests
//	@Tags		Project
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Project ID"
//	@Param		page	query		int		false	"Page"
//	@Param		size	query		int		false	"Page size"
//	@Param		status	query		string	false	"Status"	Enums(PENDING, APPROVED, REJECTED)
//	@Param		sort	query		string	false	"Sort"		Enums(id, created_at)
//	@Param		order	query		string	false	"Order"		Enums(asc, desc)
//	@Success	200		{object}	database.GetAllResponse[models.Request]
//	@Router		/project/{id}/request [get]
func GetRequests(ctx *gin.Context) {

	project := ctx.MustGet("project").(models.Project)

	var query struct {
		Status enums.RequestStatus `form:"status"`
	}

	if err := ctx.BindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	args := database.GetAllArgs[models.Request]{
		Pagination: true,
		Include: map[string][]string{
			"User":    {"Email"},
			"Project": {},
		},
		Filter: models.Request{
			ProjectID: project.ID,
			Status:    query.Status,
		},
	}

	response, err := database.GetAll(ctx.BindQuery, args)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}
