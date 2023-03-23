package main

import (
	"challange-4/routes"
)

func main() {
	var PORT = ":8080"

	routes.StartServer().Run(PORT)
}
