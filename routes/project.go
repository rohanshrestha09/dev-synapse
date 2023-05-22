package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rohanshrestha09/dev-synapse/controllers/project"
	"github.com/rohanshrestha09/dev-synapse/middleware"
	"github.com/rohanshrestha09/dev-synapse/middleware/authorize"
)

func projectRouter(router *gin.RouterGroup) {

	router.GET("/", project.GetAll)

	projectGroup := router.Group("/:id", middleware.Project())
	{
		projectGroup.GET("/", project.Get)

		projectGroup.GET("/developer", project.GetDevelopers)

		authorizedProject := projectGroup.Group("/", middleware.Auth(), authorize.Project())
		{
			authorizedProject.GET("/request", project.GetRequests)

			authorizedProject.PATCH("/", project.Update)

			authorizedProject.PATCH("/status", project.UpdateStatus)

			authorizedProject.POST("/publish", project.Publish)

			authorizedProject.DELETE("/publish", project.Unpublish)
		}
	}

	router.POST("/", middleware.Auth(), project.Create)

}
