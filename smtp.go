package main

import (
	"fmt"
	"net/smtp"
	"time"
)

// SMTP notifier
type SMTP struct {
}

// Notify the user via email
func (e SMTP) Notify(service Service) {
	t := time.Now()

	fmt.Println(t.Format("2006-01-02 15:04:05") + ": " + service.URL + " is down, sending notification via SMTP!")

	auth := smtp.PlainAuth("", SMTPUser, SMTPPassword, SMTPServer)

	to := []string{service.Options.To}
	msg := []byte("To: " + service.Options.To + "\r\n" +
		"Subject: " + replaceTokens(service, service.Options.Subject) + "\r\n" +
		"\r\n" +
		replaceTokens(service, service.Options.Body) + "\r\n")

	err := smtp.SendMail(SMTPServer+":"+SMTPPort, auth, service.Options.From, to, msg)

	if err != nil {
		fmt.Println("Error sending the notification via SMTP: " + err.Error())
	}

}
