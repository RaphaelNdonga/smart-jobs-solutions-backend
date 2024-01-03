package routes

import (
	"smartjobsolutions/controllers"
	"smartjobsolutions/middleware"
)

func SetupProviderRoutes() {
	router := GetRouter()
	router.GET("/providers/:service", middleware.AuthenticateUser, controllers.GetProviders)
	router.GET("/providers/client-posts", middleware.AuthenticateUser, controllers.GetClientPosts)
	router.POST("/providers/post", middleware.AuthenticateUser, controllers.ProviderPost)
}
