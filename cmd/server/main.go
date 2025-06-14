package main

import(
	"log"
	"net/http"

	"github.com/thesaltysleuth/notes-service/internal/api"
	"github.com/thesaltysleuth/notes-service/internal/store"

)


func main(){
	store := store.NewNoteStore()
	handler := api.NewHandler(store)
	router := api.NewRouter(handler)

	server := &http.Server{
		Addr: ":8080",
		Handler: router,
	}

	log.Println("Server starting on http://localhost:8080")
	if err:= server.ListenAndServe(); err!=nil{
		log.Fatalf("server error: %v", err)
	}
}

