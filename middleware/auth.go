package middleware

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func verifyJWT(tokenString string) (jwt.RegisteredClaims, error) {
	claims := jwt.RegisteredClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		key := []byte(os.Getenv("jwt_key"))
		return key, nil
	})
	if err != nil {
		return jwt.RegisteredClaims{}, err
	}
	log.Print("claims: ", claims)
	return claims, nil
}

func AuthenticateUser(ctx *gin.Context) {
	token := ctx.Request.Header["X-Auth-Token"]
	if token == nil {
		log.Print("X-Auth-Token nil")
		ctx.IndentedJSON(http.StatusUnauthorized, "X-Auth-Token nil ")
		ctx.Abort()
		return
	}
	claims, err := verifyJWT(token[0])
	if err != nil {
		log.Print("Error verifying jwt")
		ctx.IndentedJSON(http.StatusUnauthorized, "Error verifying jwt")
		ctx.Abort()
		return
	}
	userId := claims.Subject
	ctx.Set("userId", userId)
	ctx.Next()
}
