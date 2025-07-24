package store

import (
	"database/sql"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/thesaltysleuth/notes-service/internal/models"
)

type PGNoteStore struct { db *sql.DB }

func NewPGNoteStore(dsn string) (*PGNoteStore, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil { return nil, err }
	return &PGNoteStore{db: db},nil
}

func (s *PGNoteStore) Add(owner, title, content string) models.Note {
	id := uuid.New()
	_,_ = s.db.Exec(`INSERT INTO notes (id,owner,title,content) VALUES ($1,$2,$3,$4)`,
	id,owner,title,content)
	return models.Note{ID:id, Owner: owner, Title: title, Content: content}
}

func (s *PGNoteStore) GetAll(owner string) []models.Note {
	rows, _ := s.db.Query(`SELECT id, title, content FROM notes WHERE owner=$1`, owner)
	defer rows.Close()
	var out []models.Note
	for rows.Next() {
		var n models.Note 
		_ = rows.Scan(&n.ID, &n.Title, &n.Content)
		n.Owner = owner
		out = append(out,n)
	}
	return out
}
