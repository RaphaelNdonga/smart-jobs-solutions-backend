package controllers

import (
	"fmt"
	"log"
	"net/http"
	"smartjobsolutions/database"
	"smartjobsolutions/types"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(ctx *gin.Context) {
	userDetails := types.UserDetails{}
	if err := ctx.BindJSON(&userDetails); err != nil {
		errorMsg := "POST sign-up: bind json error"
		ctx.IndentedJSON(http.StatusInternalServerError, errorMsg)
		log.Fatal(err)
	}
	password := userDetails.Password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		errorMsg := "POST sign-up: hashing err"
		ctx.IndentedJSON(http.StatusInternalServerError, errorMsg)
		log.Fatal(err)
	}
	userDetailsDB := types.UserDetailsDB{
		Username:       userDetails.Username,
		Email:          userDetails.Email,
		HashedPassword: string(hash),
		UserType:       userDetails.UserType,
	}

	fmt.Print("success: ", userDetailsDB)
	database.AddUser(userDetailsDB)
	ctx.IndentedJSON(http.StatusOK, userDetailsDB)
}
