package main

import "fmt"

// SMTP notifier
type SMTP struct {
}

// Notify the user via email
func (e SMTP) Notify(url string) {
	fmt.Println("send email via smtp")
}
