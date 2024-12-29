package services

type NotesService struct{}

type Note struct {
	Id          string
	Name        string
	Description string
	Author      string
}

func (s *NotesService) GetNotesService() []Note {
	// Return a list of notes
	notes := []Note{
		{
			Id:          "1",
			Name:        "Note 1",
			Description: "This is the first note",
			Author:      "John Doe",
		},
		{
			Id:          "2",
			Name:        "Note 2",
			Description: "This is the second note",
			Author:      "Jane Doe",
		},
	}
	return notes
}

func (s *NotesService) CreateNoteService(data Note) Note {
	return data
}
