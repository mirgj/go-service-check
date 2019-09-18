package main

import "fmt"

// SMTP notifier
type SMTP struct {
}

// Notify the user via email
func (e SMTP) Notify(service Service) {
	fmt.Println("send email via smtp")
}
