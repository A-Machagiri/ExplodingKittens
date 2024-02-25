package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"context"
	"github.com/go-redis/redis/v8"
)

func main() {
	// Set up a simple HTTP server
	
	http.HandleFunc("/", sendHello)
	http.HandleFunc("/about", sendAbout)

	fmt.Println("Server listening on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func sendHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func sendAbout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "About page")
}

