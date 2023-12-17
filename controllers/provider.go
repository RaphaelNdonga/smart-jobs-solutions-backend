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
		log.Print("Error getting providers: ", err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, providerList)
}

func ProviderPost(ctx *gin.Context) {
	var providerPost types.PostJSON
	if err := ctx.BindJSON(&providerPost); err != nil {
		log.Print("ProviderPost Error getting post details: ", err)
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
	providerPost.Id = claims.Subject

	err = database.ProviderPost(database.GetDB(), providerPost)
	if err != nil {
		log.Print("ProviderPost Error database: ", err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, providerPost)
}
