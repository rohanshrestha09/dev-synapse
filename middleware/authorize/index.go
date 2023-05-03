package authorize

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/dev-synapse/models"
)

func Project() func(*gin.Context) {
	return func(ctx *gin.Context) {

		authUser := ctx.MustGet("auth").(models.User)

		project := ctx.MustGet("project").(models.Project)

		if authUser.ID != project.UserID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized action"})
			return
		}

		ctx.Next()

	}
}

func Request() func(*gin.Context) {
	return func(ctx *gin.Context) {

		authUser := ctx.MustGet("auth").(models.User)

		request := ctx.MustGet("request").(models.Request)

		if authUser.ID != request.UserID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized action"})
			return
		}

		ctx.Next()

	}
}
