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
		log.Print(err)
		return
	}
	clientId := ctx.GetString("userId")
	postJSON.Id = clientId

	err = database.ClientPost(database.GetDB(), postJSON)
	if err != nil {
		log.Print("ClientPost error: ", err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, postJSON)
}

func GetProviderPosts(ctx *gin.Context) {
	clientUserId := ctx.GetString("userId")

	client, err := database.GetClient(database.GetDB(), clientUserId)

	if err != nil {
		log.Print(err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	providerPosts, err := database.GetProviderPosts(database.GetDB(), client.Service)

	if err != nil {
		log.Print(err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, providerPosts)
}
