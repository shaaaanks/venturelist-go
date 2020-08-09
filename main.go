package main

import (
	"log"
	"net/http"

	"github.com/spf13/viper"

	"github.com/shaaaanks/kibisis"
)

var database kibisis.Database
var config ApplicationConfig

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error loading config file: %v", err)
	}

	err = viper.UnmarshalKey("aws", &config)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	err = validate(config)
	if err != nil {
		log.Fatalf("Configuration file error: %v", err)
	}

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
