package main

import (
	"smartjobsolutions/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	router.POST("/sign-up", controllers.PostSignUp)

	router.Run(":8000")
}
