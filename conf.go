package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Config struct {
	Admin            string
	AdminPassword    string
	DatabaseLocation string
}

func loadConfiguration() Config {
	content, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal("Error opening configuration file: ", err)
	}

	payload := Config{}
	payload.DatabaseLocation = "./database.db"
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Failed reading configuration file: ", err)
	}

	log.Printf("Admin: %s \nAdmin Password: %s\n", payload.Admin, payload.AdminPassword)
	return payload
}
