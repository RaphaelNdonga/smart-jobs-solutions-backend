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
	return router
}

func GetRouter() *gin.Engine {
	return router
}
