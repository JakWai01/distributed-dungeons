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
	sessionkey string
	password   string
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
	hitPoints   int
	sanity      int
	luck        int
	magicPoints int
	diceroll    int
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

	querySessionData(db)
	queryCharactersData(db)

}

func addSessionKey(db *sql.DB) {
	insertStatement := `
	INSERT INTO session (sessionkey, password) VALUES ('ASDE', '1234')`
	_, err := db.Exec(insertStatement)
	if err != nil {
		panic(err)
	}
}

func addCharacter(db *sql.DB) {
	insertStatement := `
	INSERT INTO characters (name, player, occupation, age, sex, residence, birthplace, str, dex, pow, con, app, edu, siz, int, hitpoints, sanity, luck, magicpoints, diceroll) VALUES ('Oliver Smith','Jakob','Doctor',36,'Male','London', 'London', 50, 60, 40, 30, 50, 50, 7, 60, 3, 15, 13, 12, 100)`
	_, err := db.Exec(insertStatement)
	if err != nil {
		panic(err)
	}
}

func updateDiceroll(db *sql.DB) {
	insertStatement := `UPDATE characters SET diceroll=1 WHERE name='Oliver Smith';`
	_, err := db.Exec(insertStatement)
	if err != nil {
		panic(err)
	}
}

func querySessionData(db *sql.DB) {
	var mySession Session
	userSQL := "SELECT sessionkey, password FROM session WHERE sessionkey = 'ASDE'"

	err := db.QueryRow(userSQL).Scan(&mySession.sessionkey, &mySession.password)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s : %s\n", mySession.sessionkey, mySession.password)
}

func queryCharactersData(db *sql.DB) {
	var myCharacters Characters
	// make WHERE condition to sessionkey with one to many relation
	userSQL := "SELECT name, player, occupation, age, sex, residence, birthplace, str, dex, pow, con, app, edu, siz, int, hitpoints, sanity, luck, magicpoints, diceroll FROM characters WHERE name='Oliver Smith'"

	err := db.QueryRow(userSQL).Scan(&myCharacters.name, &myCharacters.player, &myCharacters.occupation, &myCharacters.age, &myCharacters.sex, &myCharacters.residence, &myCharacters.birthplace, &myCharacters.str, &myCharacters.dex, &myCharacters.pow, &myCharacters.con, &myCharacters.app, &myCharacters.edu, &myCharacters.siz, &myCharacters.inte, &myCharacters.hitPoints, &myCharacters.sanity, &myCharacters.luck, &myCharacters.magicPoints, &myCharacters.diceroll)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v : %v : %v : %v\n", myCharacters.name, myCharacters.player, myCharacters.occupation, myCharacters.age)
}
