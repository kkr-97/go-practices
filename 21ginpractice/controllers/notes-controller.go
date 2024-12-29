package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/kkr-97/gin-practice/services"
)

type NotesController struct {
	notesService *services.NotesService
}

func (n *NotesController) InitNotesControllerRoutes(router *gin.Engine) {
	notesRoutes := router.Group("/notes")
	notesRoutes.GET("/", n.GetNotes())
	notesRoutes.POST("/", n.CreateNote())
}

func (n *NotesController) GetNotes() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Retrieve notes from database
		c.JSON(200, gin.H{
			"msg":   "All Notes Data retrieved successfully!!",
			"notes": n.notesService.GetNotesService(),
		})
	}
}

func (n *NotesController) CreateNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Create a new note
		c.JSON(200, gin.H{
			"msg": "Note created successfully!!",
			"note": n.notesService.CreateNoteService(services.Note{
				Id:          "3",
				Name:        "Note 3",
				Description: "This is the third note",
				Author:      "Bob Smith",
			}),
		})
	}
}
