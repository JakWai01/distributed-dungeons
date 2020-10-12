package main

import (
	"log"

	"github.com/JakWai01/distributed-dungeons/pkg/api"
)

func main() {

	api := api.NewAPI()

	if err := api.Open(); err != nil {
		log.Fatal("could not open httpServer", err)
	}

}
