package main

import (
	"smartjobsolutions/controllers"
	"smartjobsolutions/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	router := gin.Default()

	router.Use(cors.Default())

	router.POST("/sign-up", controllers.PostSignUp)

	router.Run(":8000")
}
