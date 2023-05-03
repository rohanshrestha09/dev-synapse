package routes

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.RouterGroup) {

	authRouter(router.Group("/auth"))

	userRouter(router.Group("/user"))

	ssoRouter(router.Group("/sso"))

	projectRouter(router.Group("/project"))

	requestRouter(router.Group("/request"))

}
