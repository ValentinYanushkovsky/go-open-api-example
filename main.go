package main

import (
	"log"
	"net/http"

	"hello-api/api"
	gen "hello-api/gen"
)

func main() {
	srv, err := gen.NewServer(&api.Server{})
	if err != nil {
		log.Fatalf("failed to create server: %v", err)
	}

	log.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatalf("server error: %v", err)
	}
}
