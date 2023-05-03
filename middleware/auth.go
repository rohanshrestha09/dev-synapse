package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/rohanshrestha09/dev-synapse/models"
	"github.com/rohanshrestha09/dev-synapse/service/database"
	"github.com/rohanshrestha09/dev-synapse/utils"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var jwtToken string

		bearerToken := ctx.GetHeader("Authorization")

		if strings.HasPrefix(bearerToken, "Bearer") && len(strings.Split(bearerToken, " ")) == 2 {
			jwtToken = strings.Split(bearerToken, " ")[1]
		} else {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		claims, token, err := utils.ParseJwt(jwtToken)

		if err != nil {
			if errors.Is(err, jwt.ErrSignatureInvalid) {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
				return
			}

			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		if !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		args := database.GetArgs[models.User]{
			Filter: models.User{
				Email: claims.Email,
			},
		}

		data, er := database.Get(args)

		if er != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		ctx.Set("auth", data)

		ctx.Next()

	}
}
