package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.Use(cors.Default())
	router.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "test ok",
		})
	})

	router.POST("/sign-up", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "sign up received",
		})
	})
	router.Run("localhost:8000")
}
