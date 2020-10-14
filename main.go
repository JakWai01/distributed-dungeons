package main

import (
	"log"

	"github.com/JakWai01/distributed-dungeons/pkg/api"
	"github.com/JakWai01/distributed-dungeons/pkg/database"
)

func main() {

	//api := api.NewAPI()
	//database := api.NewAPI()

	if err := database.Open(); err != nil {
		log.Fatal("Could not open database")
	}
	if err := api.Open(); err != nil {
		log.Fatal("Could not open api")
	}

}
