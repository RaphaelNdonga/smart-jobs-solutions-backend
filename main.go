package main

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type UserDetails struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	UserType string `json:"userType"`
}

func main() {
	router := gin.Default()

	router.Use(cors.Default())

	router.POST("/sign-up", func(ctx *gin.Context) {
		userDetails := UserDetails{}
		if err := ctx.BindJSON(&userDetails); err != nil {
			fmt.Println("POST sign-up: bind json error")
			return
		}
		fmt.Print("success: ", userDetails)
		ctx.IndentedJSON(http.StatusOK, userDetails)
	})
	router.Run("localhost:8000")
}
