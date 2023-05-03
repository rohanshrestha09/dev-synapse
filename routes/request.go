package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/dev-synapse/controllers/request"
	"github.com/rohanshrestha09/dev-synapse/middleware"
	"github.com/rohanshrestha09/dev-synapse/middleware/authorize"
)

func requestRouter(router *gin.RouterGroup) {

	router.Use(middleware.Auth())

	router.POST("/:id", middleware.Project(), request.Create)

	requestGroup := router.Group("/:id", middleware.Request())
	{
		authorizedRequest := requestGroup.Group("/", authorize.Request())
		{
			authorizedRequest.GET("/", request.Get)

			authorizedRequest.DELETE("/", request.Delete)

			authorizedRequest.POST("/approve", request.Approve)

			authorizedRequest.DELETE("/approve", request.Reject)
		}
	}

}
