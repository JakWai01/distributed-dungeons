package main

import (
	"encoding/json"
	"fmt"

	"log"
	"net/http"
)

// Character defines article
type Character struct {
	name        string
	player      string
	occupation  string
	age         int
	sex         string
	residence   string
	birthplace  string
	str         int
	dex         int
	pow         int
	con         int
	app         int
	edu         int
	siz         int
	inte        int
	hitpoints   int
	sanity      int
	luck        int
	magicpoints int
	diceroll    int
	sessionkey  int
}

// Characters defines array of Article
type Characters []Character

// Chars will add values to API
func Chars(w http.ResponseWriter, r *http.Request) {

	// For all characters with the same sessionkey, make a charakter in characters
	articles := Characters{
		Character{name: "Test Title", player: "Test Desc", occupation: "Test Content"},
	}

	fmt.Println("Endpoint Hit: All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}

// Change api key here on request of a new session
func handleRequests() {
	http.HandleFunc("/", Chars)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

// call queryData to get characterinformatione
// func main() {

// 	handleRequests()
// }

// NewAPI initializes API values
func NewAPI(name string, player string, occupation string, age int, sex string, residence string, birthplace string, str int, dex int, pow int, con int, app int, edu int, siz int, inte int, hitpoints int, sanity int, luck int, magicpoints int, diceroll int, sessionkey int) {
	fmt.Println(Character{name, player, occupation, age, sex, residence, birthplace, str, dex, pow, con, app, edu, siz, inte, hitpoints, sanity, luck, magicpoints, diceroll, sessionkey})

	handleRequests()
}

// Build function to set API data and call handleRequests. In Chars, then get all the char data from the matching session key and setup the api
