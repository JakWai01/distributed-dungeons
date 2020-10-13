package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// database name can be adjusted as variable
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "stripe-example"
)

// Session initializes session
type Session struct {
	sessionkey string
	password   string
}

// Characters initializes character
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

	//	addSessionKey(db)
	//	addCharacter(db)
	updateDiceroll(db)
}

func addSessionKey(db *sql.DB) {
	insertStatement := `
	INSERT INTO session (
		sessionkey, 
		password
		) 
		VALUES (
			'ASDE', 
			'1234'
			)`
	_, err := db.Exec(insertStatement)
	if err != nil {
		panic(err)
	}
	fmt.Println("Added session")
}

func addCharacter(db *sql.DB) {
	insertStatement := `
	INSERT INTO characters (
		name, 
		player, 
		occupation, 
		age, 
		sex, 
		residence, 
		birthplace, 
		str, 
		dex, 
		pow, 
		con, 
		app, 
		edu, 
		siz, 
		int, 
		hitpoints, 
		sanity, 
		luck, 
		magicpoints, 
		diceroll
		) 
		VALUES (
		'Oliver Smith',
		'Jakob',
		'Doctor',
		36,
		'Male',
		'London',
		'London',
		50,
		60,
		40,
		30,
		50,
		50,
		7,
		60,
		3,
		15,
		13,
		12,
		100
	)`
	_, err := db.Exec(insertStatement)
	if err != nil {
		panic(err)
	}
}

func getCharacter() {

}

func getSession() {

}

func updateDiceroll(db *sql.DB) {
	insertStatement := `UPDATE characters SET diceroll=1 WHERE name='Oliver Smith';`
	_, err := db.Exec(insertStatement)
	if err != nil {
		panic(err)
	}
}
