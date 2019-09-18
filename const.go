package main

import "os"

// SendGridAPIKey api for sendgrid service
var SendGridAPIKey = os.Getenv("SENDGRID_API_KEY")
