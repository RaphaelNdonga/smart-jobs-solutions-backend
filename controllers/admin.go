package controllers

import (
	"log"
	"net/http"
	"smartjobsolutions/database"
	"smartjobsolutions/types"

	"github.com/gin-gonic/gin"
)

func AddService(ctx *gin.Context) {
	var serviceJSON types.ServiceJSON
	if err := ctx.BindJSON(&serviceJSON); err != nil {
		log.Print(err)
		return
	}
	if serviceJSON.Key_Service == "" {
		log.Print("AddService error null entry")
		ctx.IndentedJSON(http.StatusBadRequest, "error null entry")
		return
	}
	err := database.AddService(database.GetDB(), serviceJSON.Key_Service)
	if err != nil {
		log.Print(err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, serviceJSON.Key_Service)
}
