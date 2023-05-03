package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/dev-synapse/controllers/auth"
	"github.com/rohanshrestha09/dev-synapse/middleware"
)

func authRouter(router *gin.RouterGroup) {

	router.Use(middleware.Auth())
	{
		router.GET("/", auth.Get)

		router.PATCH("/", auth.Update)

		router.GET("/project", auth.GetProjects)
	}

}
