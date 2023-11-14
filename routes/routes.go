package routes

import (
	"smartjobsolutions/controllers"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func SetupRouter() *gin.Engine {
	router = gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "x-auth-token"},
		ExposeHeaders:    []string{"Content-Length", "x-auth-token"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.POST("/sign-up", controllers.SignUp)
	router.POST("/sign-in", controllers.SignIn)
	router.POST("/sign-up/service-provider", controllers.RegisterServiceProvider)
	router.POST("/sign-up/client", controllers.RegisterClient)
	router.GET("/service-providers", controllers.GetServiceProviders)
	router.POST("/client/post", controllers.ClientPost)
	return router
}

func GetRouter() *gin.Engine {
	return router
}
