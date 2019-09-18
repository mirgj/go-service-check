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
	URL      string     `json:"url"`
	Name     string     `json:"name"`
	Delay    int        `json:"delay"`
	Notifier string     `json:"notifier"`
	SendGrid Sendgrid   `json:"sendgrid"`
	SMTP     SMTPStruct `json:"smtp"`
}

// SMTPStruct represent the configuration to be extended for the smtp client
type SMTPStruct struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Body    string `json:"body"`
	Subject string `json:"subject"`
}

// Sendgrid represent the sendgrid configuration
type Sendgrid struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Body    string `json:"body"`
	Subject string `json:"subject"`
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
