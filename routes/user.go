package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/dev-synapse/controllers/user"
	"github.com/rohanshrestha09/dev-synapse/middleware"
)

func userRouter(router *gin.RouterGroup) {

	router.POST("/register", user.Register)

	router.POST("/login", user.Login)

	userGroup := router.Group("/:id", middleware.User())
	{
		userGroup.GET("/", user.Get)
	}

}
