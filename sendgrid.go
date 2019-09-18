package main

import (
	"fmt"
	"time"
)

// SendGrid notifier
type SendGrid struct {
}

// Notify the user via email
func (e SendGrid) Notify(service Service) {
	t := time.Now()

	fmt.Println(t.Format("2006-01-02 15:04:05") + ": " + service.URL + " is down, sending notification via SendGrid!")

}
