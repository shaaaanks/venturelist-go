package main

import (
	"log"
	"net/http"

	"github.com/shaaaanks/kibisis"
)

var database kibisis.Database

func main() {
	var err error
	database, err = kibisis.GetDriver("arangoDB")
	if err != nil {
		log.Fatalf("Error loading database driver: %v", err)
	}

	err = database.Conn()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	err = database.Init("venturelist", "projects")
	if err != nil {
		log.Fatalf("Error initialising database: %v", err)
	}

	router := createRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
