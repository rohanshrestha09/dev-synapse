package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/dev-synapse/models"
	"github.com/rohanshrestha09/dev-synapse/service/database"
)

func Project() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		args := database.GetByIDArgs{
			Include: map[string][]string{
				"User": {"Email"},
			},
		}

		data, err := database.GetByID[models.Project](ctx.Param("id"), args)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		ctx.Set("project", data)

		ctx.Next()

	}
}
