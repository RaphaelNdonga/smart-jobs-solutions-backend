package routes

import (
	"smartjobsolutions/controllers"
	"smartjobsolutions/middleware"
)

func SetupAuthRoutes() {
	router.POST("/sign-up", controllers.SignUp)
	router.POST("/sign-in", controllers.SignIn)
	router.POST("/sign-up/provider", controllers.RegisterProvider)
	router.POST("/sign-up/client", controllers.RegisterClient)
	router.GET("/user-type", middleware.AuthenticateUser, controllers.GetUserType)
	router.GET("/userdata", middleware.AuthenticateUser, controllers.GetUserData)
	router.GET("/userposts", middleware.AuthenticateUser, controllers.GetUserPosts)
}
