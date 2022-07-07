package routes

import (
	"github.com/gin-gonic/gin"
	"go-mongo/controllers"
)

func UserRoute(router *gin.RouterGroup) {
	//	All user related route here
	router.POST("/", controllers.CreateUser)
	router.GET("/:userId", controllers.GetUser)
	router.GET("/all", controllers.GetAllUsers)
	router.PUT("/:userId", controllers.EditUser)
	router.DELETE("/:userId", controllers.DeleteUser)
}
