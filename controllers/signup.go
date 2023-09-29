package controllers

import (
	"fmt"
	"net/http"
	"smartjobsolutions/types"

	"github.com/gin-gonic/gin"
)

func PostSignUp(ctx *gin.Context) {
	userDetails := types.UserDetails{}
	if err := ctx.BindJSON(&userDetails); err != nil {
		fmt.Println("POST sign-up: bind json error")
		return
	}
	fmt.Print("success: ", userDetails)
	ctx.IndentedJSON(http.StatusOK, userDetails)
}
