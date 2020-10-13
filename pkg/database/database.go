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
	password = "mysecretpassword"
	dbname   = "distributed_dungeons"
)

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
}

// func addSessionKey(db *sql.DB) {
// 	insertStatement := `
// 	INSERT INTO sessionkeys (sessionkey, sessionpassword) VALUES
// 	(26,'rommel@500rockets.io','rommel','galisanao')`
// 	_, err = db.Exec(insertStatement)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func addCharacter(db *sql.DB) {
// 	insertStatement := `
// 	INSERT INTO personnel (p_age, p_email, p_firstname, p_lastname) VALUES
// 	(26,'rommel@500rockets.io','rommel','galisanao')`
// 	_, err = db.Exec(insertStatement)
// 	if err != nil {
// 		panic(err)
// 	}
// }
