package routes

import (
	"smartjobsolutions/controllers"
	"smartjobsolutions/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func SetupRouter() *gin.Engine {
	database.InitDB()
	router = gin.Default()
	router.Use(cors.Default())
	router.POST("/sign-up", controllers.SignUp)
	return router
}

func GetRouter() *gin.Engine {
	return router
}
