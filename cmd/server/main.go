package main

import(
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/thesaltysleuth/notes-service/internal/api"
	"github.com/thesaltysleuth/notes-service/internal/store"

)


func main(){
	noteStore := store.NewNoteStore()
	userStore := store.NewUserStore()
	handler := api.NewHandler(noteStore, userStore)
	router := api.NewRouter(handler)

	server := &http.Server{
		Addr: ":8080",
		Handler: router,
	}

	go func() {	
		log.Println("Server starting on http://localhost:8080")
		if err:= server.ListenAndServe(); err!=nil{
			log.Fatalf("server error: %v", err)
		}
	}()

	//Graceful shutdown on SIGINT
	quit := make(chan os.Signal,1)
	signal.Notify(quit, os.Interrupt)
	<-quit //wait for Ctrl+C

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err!=nil {
		log.Fatalf("graceful shutdown failed: %v", err)
	}

	log.Println("Server exited cleanly")
	
}

