package routes

import (
	"smartjobsolutions/controllers"
	"smartjobsolutions/middleware"
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
	router.POST("/sign-up/provider", controllers.RegisterProvider)
	router.POST("/sign-up/client", controllers.RegisterClient)
	router.GET("/providers/:service", middleware.AuthenticateUser, controllers.GetProviders)
	router.GET("/providers/client-posts", middleware.AuthenticateUser, controllers.GetClientPosts)
	router.GET("/user-type", middleware.AuthenticateUser, controllers.GetUserType)
	router.POST("/client/post", middleware.AuthenticateUser, controllers.ClientPost)
	router.GET("/client/provider-posts", middleware.AuthenticateUser, controllers.GetProviderPosts)
	router.POST("/providers/post", middleware.AuthenticateUser, controllers.ProviderPost)
	router.POST("/admin/add-service", controllers.AddService)
	router.GET("/get-services", controllers.GetServices)
	router.GET("/post/like/:postId", middleware.AuthenticateUser, controllers.LikePost)
	router.GET("/post/unlike/:postId", middleware.AuthenticateUser, controllers.UnlikePost)
	router.GET("/post/:postId/get-likes", middleware.AuthenticateUser, controllers.GetLikes)
	router.POST("/post/comment/:postId", middleware.AuthenticateUser, controllers.CommentPost)
	router.GET("/post/get-comments/:postId", middleware.AuthenticateUser, controllers.GetComments)
	return router
}

func GetRouter() *gin.Engine {
	return router
}
