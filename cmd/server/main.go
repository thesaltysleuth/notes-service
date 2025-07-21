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
	"github.com/thesaltysleuth/tasker"
)


func main(){
	noteStore := store.NewNoteStore()
	userStore := store.NewUserStore()

	//init Tasker (Redis @ localhost:6379 for dev)
	tq := tasker.New(os.Getenv("REDIS_ADDR"), "", 0)

	//register a task handler
	tasker.RegisterTask("index_note", func(ctx context.Context, args map[string]any) error {
		id := args["id"]
		log.Println("indexing note id", id) //pretend to push to Elastic/Typesense
		return nil
	})

	// start workers (2 goroutines)
	tq.StartWorker(context.Background(), 2)

	
	handler := api.NewHandler(noteStore, userStore, tq)
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

