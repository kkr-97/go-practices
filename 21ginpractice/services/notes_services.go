package services

import (
	"fmt"

	internal "github.com/kkr-97/gin-practice/internal/model"
	"gorm.io/gorm"
)

type NotesService struct {
	db *gorm.DB
}

func (s *NotesService) InitService(dbInstance *gorm.DB) {
	s.db = dbInstance
	s.db.AutoMigrate(&internal.Note{})

}

func (s *NotesService) GetNotesService() []internal.Note {
	// Return a list of notes
	notes := []internal.Note{}
	// Use the GORM chaining syntax to check for errors
	if err := s.db.Find(&notes).Error; err != nil {
		fmt.Println("error:", err)
		return nil
	}
	return notes
}

func (s *NotesService) CreateNoteService(note internal.Note) internal.Note {

	if err := s.db.Create(&note).Error; err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Row Inserted!!")
	}
	return note
}

func (s *NotesService) UpdateNoteService(note *internal.Note) (*internal.Note, error) {
	var noteModel *internal.Note
	if err := s.db.Where("id = ?", note.Id).First(&noteModel).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}
	// Update the note
	noteModel.Title = note.Title
	noteModel.Author = note.Author
	noteModel.Id = note.Id

	if err := s.db.Save(&noteModel).Error; err != nil {
		return nil, err
	}

	return noteModel, nil
}

func (s *NotesService) DeleteNoteService(id string) error {
	var noteModel *internal.Note
	if err := s.db.Where("id = ?", id).First(&noteModel).Error; err != nil {
		fmt.Println(err)
		return err
	}

	if err := s.db.Delete(&noteModel).Error; err != nil {
		return err
	}

	return nil
}
