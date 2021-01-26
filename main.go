package main

import (
	"modules/routes"
)

func main() {
	api.CreateApiUrl()
	// Listen and server on 0.0.0.0:8080
	api.Router.Run(":8080")
}
