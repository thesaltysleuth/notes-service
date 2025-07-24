package store

import (
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/thesaltysleuth/notes-service/internal/models"
)

type MemNoteStore struct {
	mu sync.Mutex
	notes map[string][]models.Note //key = owner
}

func NewMemNoteStore() *MemNoteStore {
	return &MemNoteStore{
		notes: make(map[string][]models.Note),
	}
}

func (s *MemNoteStore) Add(owner,title,content string) models.Note {
	s.mu.Lock()
	defer s.mu.Unlock()

	note := models.Note{
		ID:		uuid.New(),
		Title:		title,
		Content: 	content,
		CreatedAt: 	time.Now(),
		Owner: owner,
	}

	s.notes[owner] = append(s.notes[owner], note)

	return note

}


func (s *MemNoteStore) GetAll(owner string) []models.Note{
	s.mu.Lock()
	defer s.mu.Unlock()

	return append([]models.Note(nil), s.notes[owner]...) //copy
}
