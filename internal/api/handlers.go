package api

import (
	"encoding/json"
	"net/http"

	"github.com/thesaltysleuth/notes-service/internal/auth"
	"github.com/thesaltysleuth/notes-service/internal/models"
	"github.com/thesaltysleuth/notes-service/internal/store"
)

type Handler struct {
	Store *store.NoteStore
	Users *store.UserStore
}


func NewHandler(noteStore *store.NoteStore, userStore *store.UserStore) *Handler {
	return &Handler{
		Store: noteStore,
		Users: userStore,
	}
}

func (h *Handler) CreateNote(w http.ResponseWriter, r *http.Request) {
	user,ok := r.Context().Value("username").(string)
	if !ok {
		respondError(w, http.StatusUnauthorized, "unauthorized")
		return
	}
	
	var note models.Note	
	if err:= json.NewDecoder(r.Body).Decode(&note); err!=nil{
		respondError(w, http.StatusBadRequest, "invalid JSON")
		return
	}
	
	if note.Title == "" || note.Content == "" {
		respondError(w, http.StatusBadRequest, "title and content required")
		return
	}

	note.Owner = user
	note = h.Store.Add(user, note.Title, note.Content)

	respondJSON(w, http.StatusCreated, note)

}

func (h *Handler) ListNotes(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("username").(string)

	notes := h.Store.GetAll(user)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(notes)
}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	var req struct{
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err!=nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := h.Users.Add(req.Username, req.Password); err!=nil{
		http.Error(w, "user exists", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("signup success"))
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	_ = json.NewDecoder(r.Body).Decode(&req)

	if !h.Users.Validate(req.Username, req.Password) {
		http.Error(w, "invalid creds", http.StatusUnauthorized)
		return
	}

	// TODO: Issue JWT later

	token, err := auth.GenerateToken(req.Username)
	if err!= nil {
		http.Error(w, "failed to generate token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})

}

