package api

import (
	"encoding/json"
	"fmt"

	"log"
	"net/http"
)

// Character defines article
type Character struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Characters defines array of Article
type Characters []Character

// Define content here and get Information by frontend
func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Characters{
		Character{Title: "Test Title", Desc: "Test Desc", Content: "Test Content"},
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

// Main Page, unneccessary
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint Hit")
}

// Change api key here on request of a new session
func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/allArticles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
