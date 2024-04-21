package main

import (
	"log"

	"github.com/amosehiguese/restaurant-api/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}

	log.Println("Server gracefully shutdown")
}