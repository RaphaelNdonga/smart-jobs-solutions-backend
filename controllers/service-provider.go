package controllers

import (
	"log"
	"net/http"
	"smartjobsolutions/database"

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
