package main

import (
	"fmt"
	"time"
)

// Console notifier
type Console struct {
}

// Notify via console text
func (e Console) Notify(service Service) {
	t := time.Now()

	fmt.Println(t.Format("2006-01-02 15:04:05") + ": " + service.URL + " is down!")
}
