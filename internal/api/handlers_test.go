package api_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/thesaltysleuth/notes-service/internal/api"
	"github.com/thesaltysleuth/notes-service/internal/auth"
	"github.com/thesaltysleuth/notes-service/internal/store"
)


func TestCreateAndListNotes(t *testing.T) {
	nStore := store.NewNoteStore()
	uStore := store.NewUserStore()
	h := api.NewHandler(nStore,uStore)


	//fake user + token
	const user = "test"
	uStore.Add(user, "pw")
	token, _ := auth.GenerateToken(user)

	//build server
	router := api.NewRouter(h)

	// create note
	body, _ := json.Marshal(map[string]string{"title":"t", "content":"c"})
	req := httptest.NewRequest(http.MethodPost, "/notes", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusCreated {
		t.Fatalf("want 201 got %d", rec.Code)
	}

	//list notes
	req = httptest.NewRequest(http.MethodGet, "/notes", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec = httptest.NewRecorder()
	router.ServeHTTP(rec,req)

	var notes []map[string]any
	_ = json.Unmarshal(rec.Body.Bytes(), &notes)

	if len(notes) != 1 {
		t.Fatalf("want 1 note, got %d", len(notes))
	}
}


func TestAuthMissingToken(t *testing.T) {
	nStore := store.NewNoteStore()
	uStore := store.NewUserStore()
	h := api.NewHandler(nStore, uStore)
	router := api.NewRouter(h)

	//Try to access /notes without any token
	req := httptest.NewRequest(http.MethodGet, "/notes", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("expected 401, got %d", rec.Code)
	}
}

func TestInvalidMethod(t *testing.T) {
	nStore := store.NewNoteStore()
	uStore := store.NewUserStore()

	//fake user + token
	const user = "test"
	uStore.Add(user, "pw")
	token, _ := auth.GenerateToken(user)
	
	h := api.NewHandler(nStore,uStore)
	//build server
	router := api.NewRouter(h)

	// create note
	body, _ := json.Marshal(map[string]string{"title":"t", "content":"c"})
	req := httptest.NewRequest(http.MethodPut, "/notes", bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected 405, got %d", rec.Code)
	}
}
