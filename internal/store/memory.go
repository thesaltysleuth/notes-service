package store 

import (
	"sync"
	"time"

	"github.com/thesaltysleuth/notes-service/internal/models"

)

type NoteStore struct {
	mu sync.Mutex
	notes []models.Note
	nextID int
}

func NewNoteStore() *NoteStore {
	return &NoteStore{
		notes: []models.Note{},
		nextID: 1,
	}
}

func (s *NoteStore) Add(title,content string) models.Note {
	s.mu.Lock()
	defer s.mu.Unlock()

	note := models.Note{
		ID:		s.nextID,
		Title:		title,
		Content: 	content,
		CreatedAt: 	time.Now(),
	}

	s.notes = append(s.notes, note)
	s.nextID++

	return note

}


func (s *NoteStore) GetAll() []models.Note{
	s.mu.Lock()
	defer s.mu.Unlock()

	return append([]models.Note(nil), s.notes...) //copy
}
