package controllers

import (
	"log"
	"net/http"
	"smartjobsolutions/database"
	"smartjobsolutions/types"

	"github.com/gin-gonic/gin"
)

func GetServiceProviders(ctx *gin.Context) {
	serviceProviderList, err := database.GetServiceProviders(database.GetDB())
	if err != nil {
		log.Print("Error getting service providers: ", err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, serviceProviderList)
}

func ServiceProviderPost(ctx *gin.Context) {
	var serviceProviderPost types.ServiceProviderPostJSON
	if err := ctx.BindJSON(&serviceProviderPost); err != nil {
		log.Print("ServiceProviderPost Error getting post details: ", err)
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
	serviceProviderPost.Id = claims.Subject
	log.Print("claims subject: ", serviceProviderPost.Id)
	log.Print("client post json: ", serviceProviderPost)

	err = database.ServiceProviderPost(database.GetDB(), serviceProviderPost)
	if err != nil {
		log.Print("ServiceProviderPost Error database: ", err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, serviceProviderPost)
}
