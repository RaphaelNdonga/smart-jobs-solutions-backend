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

func UnlikePost(ctx *gin.Context) {
	userId := ctx.GetString("userId")
	postId := ctx.Param("postId")
	err := database.UnlikePost(database.GetDB(), userId, postId)
	if err != nil {
		log.Print(err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, "Unliked successfully")
}

func GetLikes(ctx *gin.Context) {
	postId := ctx.Param("postId")
	users, err := database.GetLikes(database.GetDB(), postId)
	if err != nil {
		log.Print(err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, users)
}

func CommentPost(ctx *gin.Context) {
	postId := ctx.Param("postId")
	userId := ctx.GetString("userId")
	type CommentPostJson struct {
		Comment string `json:"Comment"`
	}
	var commentPost CommentPostJson
	if err := ctx.BindJSON(&commentPost); err != nil {
		log.Print(err)
		return
	}
	err := database.CommentPost(database.GetDB(), postId, userId, commentPost.Comment)
	if err != nil {
		log.Print(err)
		ctx.IndentedJSON(http.StatusInternalServerError, "Commented successfully")
		return
	}
}

func GetComments(ctx *gin.Context) {
	postId := ctx.Param("postId")
	comments, err := database.GetComments(database.GetDB(), postId)
	if err != nil {
		log.Print(err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, comments)
}
