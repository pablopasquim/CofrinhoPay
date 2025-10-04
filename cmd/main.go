package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pablopasquim/CofrinhoPay/pkg/database"
)

func main() {
	database.Connect()

	router := gin.Default()
	router.Run(":8080")
}
