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
