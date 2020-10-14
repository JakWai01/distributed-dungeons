package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "stripe-example"
)

// Session initializes a session
type Session struct {
	key      int
	password string
}

// Characters initializes a character
type Characters struct {
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

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	addSessionKey(db)
	addCharacter(db)
	querySessionData(db)
	queryCharactersData(db)

}

func addSessionKey(db *sql.DB) {
	insertStatement := `
	INSERT INTO session (password) VALUES ('password123')`
	_, err := db.Exec(insertStatement)
	if err != nil {
		panic(err)
	}
}

func addCharacter(db *sql.DB) {
	insertStatement := `
	INSERT INTO character (name, player, occupation, age, sex, residence, birthplace, str, dex, pow, con, app, edu, siz, inte, hitpoints, sanity, luck, magicpoints, diceroll, sessionkey) VALUES ('Oliver Smith','Jakob','Doctor',36,'Male','London', 'London', 50, 60, 40, 30, 50, 50, 7, 60, 3, 15, 13, 12, 100, 1)`
	_, err := db.Exec(insertStatement)
	if err != nil {
		panic(err)
	}
}

func updateDiceroll(db *sql.DB) {
	insertStatement := `UPDATE characters SET dice_roll= 1 WHERE name='Oliver Smith';`
	_, err := db.Exec(insertStatement)
	if err != nil {
		panic(err)
	}
}

func querySessionData(db *sql.DB) {
	var mySession Session
	userSQL := "SELECT key, password FROM session WHERE key = 1"

	err := db.QueryRow(userSQL).Scan(&mySession.key, &mySession.password)
	if err != nil {
		panic(err)
	}

	// return array with all values
	fmt.Printf("%v : %v\n", mySession.key, mySession.password)
}

func queryCharactersData(db *sql.DB) {
	var myCharacters Characters
	// make WHERE condition to sessionkey with one to many relation
	userSQL := "SELECT name, player, occupation, age, sex, residence, birthplace, str, dex, pow, con, app, edu, siz, inte, hitpoints, sanity, luck, magicpoints, diceroll, sessionkey FROM character WHERE name='Oliver Smith'"

	err := db.QueryRow(userSQL).Scan(&myCharacters.name, &myCharacters.player, &myCharacters.occupation, &myCharacters.age, &myCharacters.sex, &myCharacters.residence, &myCharacters.birthplace, &myCharacters.str, &myCharacters.dex, &myCharacters.pow, &myCharacters.con, &myCharacters.app, &myCharacters.edu, &myCharacters.siz, &myCharacters.inte, &myCharacters.hitpoints, &myCharacters.sanity, &myCharacters.luck, &myCharacters.magicpoints, &myCharacters.diceroll, &myCharacters.sessionkey)
	if err != nil {
		panic(err)
	}
	// return array with all values
	fmt.Printf("%v : %v : %v : %v : %v\n", myCharacters.sessionkey, myCharacters.name, myCharacters.player, myCharacters.occupation, myCharacters.age)
}
