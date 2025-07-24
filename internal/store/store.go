package store

import "github.com/thesaltysleuth/notes-service/internal/models"

type NoteStore interface {
	Add(owner, title, content string) models.Note
	GetAll(owner string) []models.Note
}

