package store 

import (
	"sync"
	"time"

	"github.com/thesaltysleuth/notes-service/internal/models"

)

type NoteStore struct {
	mu sync.Mutex
	notes map[string][]models.Note //key = owner
	nextID int
}

func NewNoteStore() *NoteStore {
	return &NoteStore{
		notes: make(map[string][]models.Note),
		nextID: 1,
	}
}

func (s *NoteStore) Add(owner,title,content string) models.Note {
	s.mu.Lock()
	defer s.mu.Unlock()

	note := models.Note{
		ID:		s.nextID,
		Title:		title,
		Content: 	content,
		CreatedAt: 	time.Now(),
		Owner: owner,
	}

	s.notes[owner] = append(s.notes[owner], note)
	s.nextID++

	return note

}


func (s *NoteStore) GetAll(owner string) []models.Note{
	s.mu.Lock()
	defer s.mu.Unlock()

	return append([]models.Note(nil), s.notes[owner]...) //copy
}
