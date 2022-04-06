package main

import (
	"backend-sv/connection"
	"backend-sv/handlers"
)

func main() {
	connection.Connect()

	handlers.HandleReq()
}
