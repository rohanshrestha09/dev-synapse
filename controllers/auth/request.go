package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/dev-synapse/enums"
	"github.com/rohanshrestha09/dev-synapse/models"
	"github.com/rohanshrestha09/dev-synapse/service/database"
)

// Get Requests godoc
//
//	@Summary	Get auth requests
//	@Tags		Auth
//	@Accept		json
//	@Produce	json
//	@Param		page	query		int		false	"Page"
//	@Param		size	query		int		false	"Page size"
//	@Param		status	query		string	false	"Status"	Enums(PENDING, APPROVED, REJECTED)
//	@Param		sort	query		string	false	"Sort"		Enums(id, created_at)
//	@Param		order	query		string	false	"Order"		Enums(asc, desc)
//	@Success	200		{object}	database.GetAllResponse[models.Request]
//	@Router		/auth/request [get]
func GetRequests(ctx *gin.Context) {

	authUser := ctx.MustGet("auth").(models.User)

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
			"Project": {},
		},
		Filter: models.Request{
			UserID: authUser.ID,
			Status: query.Status,
		},
	}

	response, err := database.GetAll(ctx.BindQuery, args)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}
