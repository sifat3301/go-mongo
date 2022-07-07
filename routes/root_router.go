package routes

import "github.com/gin-gonic/gin"

func RootRoute(router *gin.Engine) {

	apiRoute := router.Group("api/v1")
	userRoute := apiRoute.Group("/user")
	UserRoute(userRoute)
}
