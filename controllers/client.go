package controllers

import (
	"log"
	"net/http"
	"smartjobsolutions/database"
	"smartjobsolutions/types"

	"github.com/gin-gonic/gin"
)

func ClientPost(ctx *gin.Context) {
	var postJSON types.PostJSON
	err := ctx.BindJSON(&postJSON)
	if err != nil {
		log.Print("ClientPost error: ", err)
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	token := ctx.Request.Header["X-Auth-Token"]
	if token == nil {
		log.Print("ClientPost error: token is nil")
		ctx.IndentedJSON(http.StatusUnauthorized, token)
		return
	}
	claims, err := VerifyJWT(token[0])
	if err != nil {
		log.Print("Client post error: ", err)
	}
	postJSON.Id = claims.Subject
	log.Print("claims subject: ", postJSON.Id)
	log.Print("client post json: ", postJSON)

	err = database.ClientPost(database.GetDB(), postJSON)
	if err != nil {
		log.Print("ClientPost error: ", err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, postJSON)
}

func GetClientPosts(ctx *gin.Context) {
	clientPosts, err := database.GetClientPosts(database.GetDB())
	if err != nil {
		log.Print("GetClientPosts error from database: ", err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, clientPosts)
}
