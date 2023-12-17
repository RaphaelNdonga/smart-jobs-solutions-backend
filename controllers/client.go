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
		ctx.IndentedJSON(http.StatusBadRequest, err)
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
