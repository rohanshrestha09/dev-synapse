package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/dev-synapse/enums"
	"github.com/rohanshrestha09/dev-synapse/models"
	"github.com/rohanshrestha09/dev-synapse/service/database"
)

// Get Projects godoc
//
//	@Summary	Get all projects
//	@Tags		Auth
//	@Accept		json
//	@Produce	json
//	@Param		page		query		int		false	"Page"
//	@Param		size		query		int		false	"Page size"
//	@Param		status		query		string	false	"Status"	Enums(OPEN, IN_PROGRESS, COMPLETED)
//	@Param		published	query		boolean	false	"Published"
//	@Param		sort		query		string	false	"Sort"	Enums(id, created_at, name)
//	@Param		order		query		string	false	"Order"	Enums(asc, desc)
//	@Param		search		query		string	false	"Search"
//	@Success	200			{object}	database.GetAllResponse[models.Project]
//	@Router		/auth/project [get]
//	@Security	Bearer
func GetProjects(ctx *gin.Context) {

	authUser := ctx.MustGet("auth").(models.User)

	var query struct {
		Published bool                `form:"published"`
		Status    enums.ProjectStatus `form:"status"`
	}

	if err := ctx.BindQuery(&query); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var mapFilter = make(map[string]any)

	if ctx.Query("published") != "" {
		mapFilter["Published"] = query.Published
	}

	args := database.GetAllArgs[models.Project]{
		Pagination: true,
		Search:     true,
		Filter: models.Project{
			UserID: authUser.ID,
			Status: query.Status,
		},
		MapFilter: mapFilter,
	}

	response, err := database.GetAll(ctx.BindQuery, args)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
