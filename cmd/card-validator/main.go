package main

import (
	"log"
	"net/http"
	"test-ms-credit-card-validator/handlers"
)

func main() {
	http.HandleFunc("/validate", handlers.ValidateCardHandler)

	log.Println("Starting server on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
