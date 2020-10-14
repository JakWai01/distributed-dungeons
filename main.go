package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/JakWai01/distributed-dungeons/pkg/handlers"
	_ "github.com/lib/pq"
)

func main() {
	laddr := flag.String("laddr", ":8080", "Listen address")
	dbHost := flag.String("dbHost", "localhost", "DB host")
	dbPort := flag.Int("dbPort", 5432, "DB port")
	dbUser := flag.String("dbUser", "postgres", "DB user")
	dbPassword := flag.String("dbPassword", "postgres", "DB password")
	dbName := flag.String("dbName", "distributed_dungeons", "DB database name")

	dbConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		*dbHost, *dbPort, *dbUser, *dbPassword, *dbName)
	db, err := sql.Open("postgres", dbConn)
	if err != nil {
		log.Fatal("could not connect to db", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("could not ping db", err)
	}

	log.Printf("connected to db %v\n", *dbHost)

	http.HandleFunc("/characters", handlers.GetAllCharactersHandler(db))

	log.Printf("listening on %v\n", *laddr)

	log.Fatal(http.ListenAndServe(*laddr, nil))
}
