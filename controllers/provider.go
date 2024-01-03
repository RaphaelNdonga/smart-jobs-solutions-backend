package controllers

import (
	"log"
	"net/http"
	"smartjobsolutions/database"
	"smartjobsolutions/types"

	"github.com/gin-gonic/gin"
)

func GetProviders(ctx *gin.Context) {
	service := ctx.Param("service")
	providerList, err := database.GetProviders(database.GetDB(), service)
	if err != nil {
		log.Print(err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, providerList)
}

func ProviderPost(ctx *gin.Context) {
	var providerPost types.PostJSON
	if err := ctx.BindJSON(&providerPost); err != nil {
		log.Print(err)
		return
	}
	userId := ctx.GetString("userId")
	log.Print("provider userid: ", userId)
	providerPost.Id = userId

	err := database.ProviderPost(database.GetDB(), providerPost)
	if err != nil {
		log.Print(err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, providerPost)
}

func GetClientPosts(ctx *gin.Context) {
	providerUserId := ctx.GetString("userId")

	provider, err := database.GetProvider(database.GetDB(), providerUserId)

	if err != nil {
		log.Print(err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	clientPosts, err := database.GetClientPosts(database.GetDB(), provider.Service)

	if err != nil {
		log.Print(err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, clientPosts)
}
