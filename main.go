package main

import (
	"smartjobsolutions/database"
	"smartjobsolutions/routes"
)

func main() {
	database.InitDB()
	router := routes.SetupRouter()

	router.Run(":8000")
}
