package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/start", start)
	http.HandleFunc("/move", move)
	http.HandleFunc("/end", end)
	http.HandleFunc("/", index)
	port := os.Getenv("PORT")
	if port == "" {
		port = "9001"
	}

	log.Printf("Running server on port %s...\n", port)
	http.ListenAndServe(":"+port, nil)
}
