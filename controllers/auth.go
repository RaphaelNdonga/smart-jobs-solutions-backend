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
		log.Print("SignUp Error binding json: ", err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	password := userDetails.Password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		log.Print("SignUp Error generating password: ", err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	userDetailsDB := types.UserDetailsDB{
		Username:       userDetails.Username,
		Email:          userDetails.Email,
		HashedPassword: string(hash),
		UserType:       userDetails.UserType,
	}

	fmt.Print("success: ", userDetailsDB)
	db := database.GetDB()
	database.AddUser(db, userDetailsDB)
	ctx.IndentedJSON(http.StatusOK, userDetailsDB)
}

func SignIn(ctx *gin.Context) {
	userDetails := types.UserDetails{}

	if err := ctx.BindJSON(&userDetails); err != nil {
		log.Print("SignIn Error binding json: ", err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	password := []byte(userDetails.Password)
	db := database.GetDB()
	dbUser, err := database.GetUserByEmail(db, userDetails.Email)

	if err != nil {
		log.Print("SignIn Error from db: ", err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.HashedPassword), password)
	if err != nil {
		log.Print("SignIn Error comparing hash and password", err)
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, dbUser)
}
