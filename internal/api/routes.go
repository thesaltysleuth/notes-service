package api

import (
	"fmt"
	"net/http"
)

func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from structured Go")
	})

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintln(w, "I am health")
	})

	return mux
}
