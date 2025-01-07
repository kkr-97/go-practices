package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	internal "github.com/kkr-97/gin-practice/internal/model"
	"github.com/kkr-97/gin-practice/services"
)

type NotesController struct {
	notesService *services.NotesService
}

func (n *NotesController) InitControllers(notesService *services.NotesService) {
	n.notesService = notesService
}

func (n *NotesController) InitNotesRoutes(router *gin.Engine) {
	notesRoutes := router.Group("/notes")
	notesRoutes.GET("/", n.GetNotes())
	notesRoutes.POST("/", n.CreateNote())
	notesRoutes.PUT("/", n.UpdateNote())
	notesRoutes.DELETE("/:id", n.DeleteNote())
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
		var note internal.Note
		if err := c.BindJSON(&note); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid request body",
			})
			return
		}
		// Create a new note in the database
		note = n.notesService.CreateNoteService(note)
		// Return the newly created note
		c.JSON(201, gin.H{
			"msg":  "Note created successfully!!",
			"note": note,
		})
	}
}

func (n *NotesController) UpdateNote() gin.HandlerFunc {
	return func(c *gin.Context) {
		var note *internal.Note

		if err := c.BindJSON(&note); err != nil {
			c.JSON(400, gin.H{
				"error": "Invalid request body",
			})
			return
		}
		// Create a new note in the database
		note, err := n.notesService.UpdateNoteService(note)
		if err != nil {
			val := fmt.Sprintf("%v", err)
			c.JSON(500, gin.H{
				"error": val,
			})
			return
		}
		// Return the newly created note
		c.JSON(201, gin.H{
			"msg":  "Note updated successfully!!",
			"note": note,
		})
	}
}

func (n *NotesController) DeleteNote() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")
		err := n.notesService.DeleteNoteService(id)
		if err != nil {
			val := fmt.Sprintf("%v", err)
			c.JSON(500, gin.H{
				"error": val,
			})
			return
		}
		// Return the newly created note
		c.JSON(201, gin.H{
			"msg": "Note deleted successfully!!",
			"id":  id,
		})
	}
}
