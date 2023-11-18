package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header["X-Auth-Token"]
	if token == nil {
		ctx.IndentedJSON(http.StatusUnauthorized, "X-Auth-Token nil")
		ctx.Abort()
		return
	}
	ctx.Next()
}
