package routes

import "smartjobsolutions/controllers"

func SetupAuthRoutes() {
	router.POST("/sign-up", controllers.SignUp)
	router.POST("/sign-in", controllers.SignIn)
	router.POST("/sign-up/provider", controllers.RegisterProvider)
	router.POST("/sign-up/client", controllers.RegisterClient)
}
