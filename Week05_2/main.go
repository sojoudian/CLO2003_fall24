package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

type Message struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Welcome to the Go API\n")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	var msg Message
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid read request body", http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &msg)
	if err != nil {
		http.Error(w, "Invalid read request body", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func main() {
	// Define a flag for the port number with a default value of 8091
	port := flag.String("port", "8091", "Port number for the server")
	flag.Parse()

	// Use the port from user input or the default
	pNumber := fmt.Sprintf(":%s", *port)
	fmt.Printf("Server is running on the port: %s\n", pNumber)

	// Register the handlers
	http.HandleFunc("/", getHandler)
	http.HandleFunc("/post", postHandler)

	// Start the server
	err := http.ListenAndServe(pNumber, nil)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
