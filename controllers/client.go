package controllers

import (
	"log"
	"net/http"
	"smartjobsolutions/database"
	"smartjobsolutions/types"

	"github.com/gin-gonic/gin"
)

func ClientPost(ctx *gin.Context) {
	var clientPostJSON types.ClientPostJSON
	err := ctx.BindJSON(&clientPostJSON)
	if err != nil {
		log.Print("ClientPost error: ", err)
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	clientPostResponse, err := database.ClientPost(database.GetDB(), clientPostJSON)
	if err != nil {
		log.Print("ClientPost error: ", err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, clientPostResponse)
}
