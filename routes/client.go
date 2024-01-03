package routes

import (
	"smartjobsolutions/controllers"
	"smartjobsolutions/middleware"
)

func SetupClientRoutes() {
	router := GetRouter()
	router.POST("/client/post", middleware.AuthenticateUser, controllers.ClientPost)
	router.GET("/client/provider-posts", middleware.AuthenticateUser, controllers.GetProviderPosts)
}
