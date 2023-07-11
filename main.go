package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Struct untuk data Response
type Response struct {
	Message string `json:"message"`
}

func main() {
	// Endpoint untuk API
	http.HandleFunc("/api/hello", helloHandler)

	// Menjalankan server di port 8080
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Handler untuk route /api/hello
func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Mengatur header response sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// Membuat response
	response := Response{Message: "Hello, World!"}

	// Mengirim response dalam format JSON
	json.NewEncoder(w).Encode(response)
}
