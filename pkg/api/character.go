package main

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

// Chars will add values to API
func Chars(w http.ResponseWriter, r *http.Request) {
	articles := Characters{
		Character{Title: "Test Title", Desc: "Test Desc", Content: "Test Content"},
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

// Change api key here on request of a new session
func handleRequests() {
	http.HandleFunc("/", Chars)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}
