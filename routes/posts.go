package routes

import (
	"smartjobsolutions/controllers"
	"smartjobsolutions/middleware"
)

func SetupPostsRoutes() {
	router := GetRouter()
	router.GET("/post/like/:postId", middleware.AuthenticateUser, controllers.LikePost)
	router.GET("/post/unlike/:postId", middleware.AuthenticateUser, controllers.UnlikePost)
	router.GET("/post/:postId/get-likes", middleware.AuthenticateUser, controllers.GetLikes)
	router.POST("/post/comment/:postId", middleware.AuthenticateUser, controllers.CommentPost)
	router.GET("/post/get-comments/:postId", middleware.AuthenticateUser, controllers.GetComments)
}
