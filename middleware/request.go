package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/dev-synapse/models"
	"github.com/rohanshrestha09/dev-synapse/service/database"
)

func Request() func(ctx *gin.Context) {
	return func(ctx *gin.Context) {

		args := database.GetByIDArgs{
			Include: map[string][]string{
				"User":    {"Email"},
				"Project": {},
			},
		}

		data, err := database.GetByID[models.Request](ctx.Param("id"), args)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		ctx.Set("request", data)

		ctx.Next()

	}
}
