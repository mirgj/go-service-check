package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Services is a collection of service
type Services struct {
	Services []Service `json:"sites"`
}

// Service is a single service configuration with the defined property
type Service struct {
	URL      string `json:"url"`
	Delay    int    `json:"delay"`
	Notifier string `json:"notifier"`
}

func readFile() Services {
	jsonFile, err := os.Open("db.json")

	if err != nil {
		printErrorAndExit("Not able to read the file: db.json", err)
	}

	byteValue, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		printErrorAndExit("Error reading the JSON file", err)
	}

	var sites Services
	err = json.Unmarshal(byteValue, &sites)

	if err != nil {
		printErrorAndExit("Invalid JSON structure", err)
	}

	return sites
}

func printErrorAndExit(message string, err error) {
	fmt.Println(message)
	fmt.Println(err)

	os.Exit(1)
}
