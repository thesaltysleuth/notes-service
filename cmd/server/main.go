package main

import(
	"log"
	"net/http"

	"github.com/thesaltysleuth/notes-service/internal/api"
)


func main(){
	handler := api.NewRouter()

	server := &http.Server{
		Addr: ":8080",
		Handler: handler,
	}

	log.Println("Server starting on http://localhost:8080")
	if err:= server.ListenAndServe(); err!=nil{
		log.Fatalf("server error: %v", err)
	}
}

