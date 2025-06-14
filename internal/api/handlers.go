package api

import (
	"encoding/json"
	"net/http"

	"github.com/thesaltysleuth/notes-service/internal/store"

)

type Handler struct {
	Store *store.NoteStore
}


func NewHandler(store *store.NoteStore) *Handler {
	return &Handler{Store: store}
}

func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	var req struct{
		Title	string	`json:"title"`
		Content	string 	`json:"content"`
	}
	if err:= json.NewDecoder(r.Body).Decode(&req); err!=nil{
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	note := h.Store.Add(req.Title, req.Content)

	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(note)

}

func (h *Handler) ListNotes(w http.ResponseWriter, r *http.Request) {
	notes := h.Store.GetAll()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

