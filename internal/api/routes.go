package api

import (
	"fmt"
	"net/http"
)

func NewRouter(h *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/notes", AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		switch r.Method {
		case http.MethodGet:
			h.ListNotes(w,r)
		case http.MethodPost:
			h.CreateNote(w,r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})))

	mux.HandleFunc("/signup", h.Signup)
	mux.HandleFunc("/login", h.Login)


	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "I am health")
	})

	return LoggingMiddleware(mux)
}
