package models

import (
	"time"

	"github.com/google/uuid"
)


type Note struct{
	ID		uuid.UUID		`json:"id"`
	Title		string		`json:"title"`
	Content		string		`json:"content"`
	CreatedAt	time.Time 	`json:"created_at"`
	Owner 		string `json:"-"` // Hide from JSON response for now
}


