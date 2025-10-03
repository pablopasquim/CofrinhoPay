package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pablopasquim/CofrinhoPay/internal/database"
)

func main() {

	r := gin.Default() // create router/engine API default.

	database.Connect() // connect db
	log.Println("Server on port http://localhost:8080 (default)")
	err := r.Run(":8080") // run the server
	if err != nil {
		return
	}

	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)

}
