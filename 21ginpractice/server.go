package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func serveHome(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "Welcome to the GIN!!",
	})
}

func main() {
	server := gin.Default()

	server.GET("/", serveHome)

	server.Run(":8080")
	fmt.Println("Server is running on 8080 port!!")
}
