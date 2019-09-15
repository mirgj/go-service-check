package main

import "fmt"

// SendGrid notifier
type SendGrid struct {
}

// Notify the user via email
func (e SendGrid) Notify(url string) {
	fmt.Println("send sendgrid")
}
