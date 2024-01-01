package controllers

import (
	"log"
	"net/http"
	"smartjobsolutions/database"

	"github.com/gin-gonic/gin"
)

func LikePost(ctx *gin.Context) {
	userId := ctx.GetString("userId")
	postId := ctx.Param("postId")
	err := database.LikePost(database.GetDB(), userId, postId)
	if err != nil {
		log.Print(err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, "Liked successfully")
}
