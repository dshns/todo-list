package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

func main() {
	port := ":7540"

	if envPort, exists := os.LookupEnv("TODO_PORT"); exists {
		port = fmt.Sprintf(":%s", envPort)
	}

	router := chi.NewRouter()

	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.Dir("web"))))

	log.Printf("Server started on http://localhost%v", port)

	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Server startup error: %v", err)
	}
}
