package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

// In-memory store for mapped URLs and a mutex to handle concurrent requests safely
var (
	urlStore = make(map[string]string)
	mutex    = &sync.Mutex{}
)

// Generates a random 6-character short code
func generateShortCode() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Handler to convert a long URL into a short URL
func shortenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	var data struct {
		URL string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil || data.URL == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	code := generateShortCode()

	mutex.Lock()
	urlStore[code] = data.URL
	mutex.Unlock()

	shortURL := fmt.Sprintf("http://localhost:8080/%s", code)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"short_url": shortURL,
		"code":      code,
	})
}

// Handler to redirect from a short code to the original URL
func redirectHandler(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Path[1:] // Extract short code from URL path (e.g., extracts "abc123" from "/abc123")

	mutex.Lock()
	originalURL, exists := urlStore[code]
	mutex.Unlock()

	if !exists {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound) // 302 Found redirect
}

func main() {
	http.HandleFunc("/shorten", shortenHandler)
	http.HandleFunc("/", redirectHandler)

	// Note: Server listens on port 8080 (fixed port mismatch in print message)
	fmt.Println("URL Shortener running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}