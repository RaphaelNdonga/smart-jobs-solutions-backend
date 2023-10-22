package main

import (
	"smartjobsolutions/routes"
)

func main() {

	router := routes.SetupRouter()

	router.Run(":8000")
}
