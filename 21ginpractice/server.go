package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kkr-97/gin-practice/controllers"
	internal "github.com/kkr-97/gin-practice/internal/database"
	"github.com/kkr-97/gin-practice/services"
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
	newId := ctx.Param("newId")

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
	db := internal.InitDB()

	if db == nil {
		log.Fatal("Failed to connect to the database")
	}

	noteService := &services.NotesService{}
	noteService.InitService(db)

	server.GET("/", serveHome)

	//working on params
	// server.GET("/:id/:newId", changeId)

	// //binding or consuming JSON data into struct
	// server.POST("/new_user", userDetails)

	notesController := &controllers.NotesController{}
	notesController.InitNotesRoutes(server)
	notesController.InitControllers(noteService)

	server.Run(":8080")
	fmt.Println("Server is running on 8080 port!!")
}
