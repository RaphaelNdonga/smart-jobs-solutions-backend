package routes

import (
	"smartjobsolutions/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func SetupRouter() *gin.Engine {
	router = gin.Default()
	router.Use(cors.Default())
	router.POST("/sign-up", controllers.SignUp)
	router.POST("/sign-in", controllers.SignIn)
	router.POST("/sign-up/service-provider", controllers.RegisterServiceProvider)
	router.POST("/sign-up/client", controllers.RegisterClient)
	router.GET("/service-providers", controllers.GetServiceProviders)
	return router
}

func GetRouter() *gin.Engine {
	return router
}
