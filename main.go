package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/viper"

	"github.com/shaaaanks/kibisis"
)

var database kibisis.Database
var config ApplicationConfig

func initialiseConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("Error loading config file: %v", err)
	}

	err = viper.UnmarshalKey("aws", &config)
	if err != nil {
		return fmt.Errorf("Error reading config file - aws: %v", err)
	}

	err = viper.UnmarshalKey("database", &config)
	if err != nil {
		return fmt.Errorf("Error reading config file - database: %v", err)
	}

	err = validate(config)
	if err != nil {
		return fmt.Errorf("Configuration file error: %v", err)
	}

	return nil
}

func initialiseDatabase() error {
	var err error
	database, err = kibisis.GetDriver(config.DatabaseDriver)
	if err != nil {
		return fmt.Errorf("Error loading database driver: %v", err)
	}

	err = database.Conn(config.DatabaseHost, config.DatabaseUsername, config.DatabasePassword)
	if err != nil {
		return fmt.Errorf("Error connecting to database: %v", err)
	}

	err = database.Init(config.Database, config.DatabaseCollection)
	if err != nil {
		return fmt.Errorf("Error initialising database: %v", err)
	}

	return nil
}

func main() {
	err := initialiseConfig()
	if err != nil {
		log.Fatal(err)
	}

	err = initialiseDatabase()
	if err != nil {
		log.Fatal(err)
	}

	router := createRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
