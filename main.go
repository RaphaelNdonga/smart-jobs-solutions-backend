package main

import (
	"log"
	"smartjobsolutions/database"
	"smartjobsolutions/routes"

	"github.com/joho/godotenv"
)

func main() {
	database.InitDB()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Main error loading dotenv: ", err)
	}

	router := routes.SetupRouter()

	router.Run(":8000")
}
