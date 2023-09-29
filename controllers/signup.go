package controllers

import (
	"fmt"
	"net/http"
	"smartjobsolutions/types"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func PostSignUp(ctx *gin.Context) {
	userDetails := types.UserDetails{}
	if err := ctx.BindJSON(&userDetails); err != nil {
		errorMsg := "POST sign-up: bind json error"
		fmt.Println(errorMsg)
		ctx.IndentedJSON(http.StatusInternalServerError, errorMsg)
		return
	}
	password := userDetails.Password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		errorMsg := "POST sign-up: hashing err"
		fmt.Println(errorMsg)
		ctx.IndentedJSON(http.StatusInternalServerError, errorMsg)
		return
	}
	userDetailsDB := types.UserDetailsDB{
		Username:       userDetails.Username,
		Email:          userDetails.Email,
		HashedPassword: hash,
		UserType:       userDetails.UserType,
	}

	fmt.Print("success: ", userDetailsDB)
	ctx.IndentedJSON(http.StatusOK, userDetailsDB)
}
