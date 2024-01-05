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
		log.Print(err)
		return
	}
	password := userDetails.Password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	if err != nil {
		log.Print(err)
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
		log.Print(err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	jwtToken, err := generateJWT(uuid)
	log.Print("jwt-token: ", jwtToken)
	if err != nil {
		log.Print(err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.Header("access-control-expose-headers", "x-auth-token")
	ctx.Header("x-auth-token", jwtToken)
	ctx.IndentedJSON(http.StatusOK, uuid)
}

func RegisterProvider(ctx *gin.Context) {
	providerJSON := types.ProviderJSON{}

	if err := ctx.BindJSON(&providerJSON); err != nil {
		log.Print(err)
		return
	}
	provider := types.Provider(providerJSON)

	if err := database.AddProvider(database.GetDB(), provider); err != nil {
		log.Print(err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, provider)
}

func RegisterClient(ctx *gin.Context) {
	clientJSON := types.ClientJSON{}
	err := ctx.BindJSON(&clientJSON)
	if err != nil {
		log.Print(err)
		return
	}
	client := types.Client(clientJSON)
	err = database.AddClient(database.GetDB(), client)
	if err != nil {
		log.Print(err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, client)
}

func SignIn(ctx *gin.Context) {
	userDetails := types.UserDetails{}

	if err := ctx.BindJSON(&userDetails); err != nil {
		log.Print(err)
		return
	}
	password := []byte(userDetails.Password)
	db := database.GetDB()
	dbUser, err := database.GetUserByEmail(db, userDetails.Email)

	if err != nil {
		log.Print(err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.HashedPassword), password)
	if err != nil {
		log.Print(err)
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	jwtToken, err := generateJWT(dbUser.Id)
	if err != nil {
		log.Print(err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.Header("access-control-expose-headers", "x-auth-token")
	ctx.Header("x-auth-token", jwtToken)
	ctx.IndentedJSON(http.StatusOK, dbUser)
}

func GetUserType(ctx *gin.Context) {
	userId := ctx.GetString("userId")
	userdetails, err := database.GetUserById(database.GetDB(), userId)
	if err != nil {
		log.Print(err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, userdetails.UserType)
}

func GetServices(ctx *gin.Context) {
	services, err := database.GetServices(database.GetDB())
	if err != nil {
		log.Print(err)
		ctx.IndentedJSON(http.StatusBadGateway, err)
		return
	}
	ctx.IndentedJSON(http.StatusOK, services)
}

func GetUserData(ctx *gin.Context) {
	userId := ctx.GetString("userId")
	userdetails, err := database.GetUserById(database.GetDB(), userId)
	if err != nil {
		log.Print(err)
		ctx.IndentedJSON(http.StatusInternalServerError, err)
		return
	}
	type Profile struct {
		Username string
		Location string
	}
	profile := Profile{
		Location: userdetails.Location,
		Username: userdetails.Username,
	}

	ctx.IndentedJSON(http.StatusOK, profile)

}
