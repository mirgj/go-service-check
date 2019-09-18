package main

import (
	"os"
	"strconv"
	"strings"
)

var (
	// SendGridAPIKey api for sendgrid service
	SendGridAPIKey = os.Getenv("SENDGRID_API_KEY")
	// SMTPServer SMTP server host
	SMTPServer = os.Getenv("SMTP_SERVER")
	// SMTPUser SMTP username (optional only for auth)
	SMTPUser = os.Getenv("SMTP_USER")
	// SMTPPassword password of the SMTP service (optional)
	SMTPPassword = os.Getenv("SMTP_PASSWORD")
	// SMTPPort port of the SMTP service
	SMTPPort = os.Getenv("SMTP_PORT")
)

func replaceTokens(service Service, str string) string {
	newStr := strings.Replace(str, "{{url}}", service.URL, -1)
	newStr = strings.Replace(newStr, "{{name}}", service.Name, -1)
	newStr = strings.Replace(newStr, "{{delay}}", strconv.Itoa(service.Delay), -1)

	return newStr
}
