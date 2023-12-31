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
	SetupPostsRoutes()
	SetupClientRoutes()
	SetupProviderRoutes()
	SetupAuthRoutes()
	router.POST("/admin/add-service", controllers.AddService)
	router.GET("/get-services", controllers.GetServices)
	return router
}

func GetRouter() *gin.Engine {
	return router
}
