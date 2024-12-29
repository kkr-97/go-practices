package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kkr-97/gin-practice/controllers"
)

func serveHome(ctx *gin.Context) {
	// case1: sending response using JSON method
	ctx.JSON(200, gin.H{
		"message": "Welcome to the GIN!!",
	})

	//case2: sending response using String method
	// ctx.String(200, "Welcome to the GIN!!")

	//case3: using HTML method
	// ctx.HTML(200, "index.html", gin.H{
	// 	"title":   "Welcome",
	// 	"content": "This is a dynamic page!",
	// })
}

func changeId(ctx *gin.Context) {
	id := ctx.Param("id")
	newId := ctx.Param(("newId"))

	ctx.JSON(200, gin.H{
		"message": "ID changed successfully",
		"newId":   newId,
		"oldId":   id,
	})

}

func userDetails(ctx *gin.Context) {
	type User struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var newUser User

	ctx.BindJSON(&newUser)

	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "User created successfully",
		"email": newUser.Email,
	})
}

func main() {
	server := gin.Default()

	server.GET("/", serveHome)

	// //working on params
	// server.GET("/:id/:newId", changeId)

	// //binding or consuming JSON data into struct
	// server.POST("/new_user", userDetails)

	notesController := &controllers.NotesController{}
	notesController.InitNotesControllerRoutes(server)

	server.Run(":8080")
	fmt.Println("Server is running on 8080 port!!")
}
