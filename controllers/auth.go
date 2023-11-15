package controllers

import (
	"fmt"
	"log"
	"net/http"
	"smartjobsolutions/database"
	"smartjobsolutions/types"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func generateJWT(uuid string) (string, error) {
	var (
		key          []byte
		token        *jwt.Token
		signedString string
	)

	key = []byte(os.Getenv("jwt_key"))
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "smartjobsolutions-server",
		"sub": uuid,
	})
	signedString, err := token.SignedString(key)

	return signedString, err
}

func VerifyJWT(tokenString string) (jwt.RegisteredClaims, error) {
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

func SignUp(ctx *gin.Context) {
	userDetails := types.UserDetails{}
	if err := ctx.BindJSON(&userDetails); err != nil {
		log.Print("SignUp Error binding json: ", err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	password := userDetails.Password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		log.Print("SignUp Error generating password: ", err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	userDetailsDB := types.UserDetailsDB{
		Username:       userDetails.Username,
		Email:          userDetails.Email,
		HashedPassword: string(hash),
		Location:       userDetails.Location,
		UserType:       userDetails.UserType,
	}

	fmt.Print("success: ", userDetailsDB)
	db := database.GetDB()
	uuid, err := database.AddUser(db, userDetailsDB)
	if err != nil {
		log.Print("SignUp Error fetching user id: ", err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	jwtToken, err := generateJWT(uuid)
	log.Print("jwt-token: ", jwtToken)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.Header("access-control-expose-headers", "x-auth-token")
	ctx.Header("x-auth-token", jwtToken)
	ctx.IndentedJSON(http.StatusOK, uuid)
}

func RegisterServiceProvider(ctx *gin.Context) {
	serviceProviderJSON := types.ServiceProviderJSON{}

	if err := ctx.BindJSON(&serviceProviderJSON); err != nil {
		log.Print("RegisterServiceProvider json error: ", err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	serviceProvider := types.ServiceProvider(serviceProviderJSON)

	if err := database.AddServiceProvider(database.GetDB(), serviceProvider); err != nil {
		log.Print("RegisterServiceProvider database Error: ", err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, serviceProvider)
}

func RegisterClient(ctx *gin.Context) {
	clientJSON := types.ClientJSON{}
	err := ctx.BindJSON(&clientJSON)
	if err != nil {
		log.Print("RegisterClient json error: ", err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	client := types.Client(clientJSON)
	err = database.AddClient(database.GetDB(), client)
	if err != nil {
		log.Print("RegisterClient database error: ", err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, client)
}

func SignIn(ctx *gin.Context) {
	userDetails := types.UserDetails{}

	if err := ctx.BindJSON(&userDetails); err != nil {
		log.Print("SignIn Error binding json: ", err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	password := []byte(userDetails.Password)
	db := database.GetDB()
	dbUser, err := database.GetUserByEmail(db, userDetails.Email)

	if err != nil {
		log.Print("SignIn Error from db: ", err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.HashedPassword), password)
	if err != nil {
		log.Print("SignIn Error comparing hash and password", err)
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	jwtToken, err := generateJWT(userDetails.Username)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.Header("access-control-expose-headers", "x-auth-token")
	ctx.Header("x-auth-token", jwtToken)
	ctx.IndentedJSON(http.StatusOK, dbUser)
}
