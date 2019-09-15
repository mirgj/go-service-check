package main

import (
	"fmt"
	"time"
)

// Console notifier
type Console struct {
}

// Notify via console text
func (e Console) Notify(url string) {
	t := time.Now()

	fmt.Println(t.Format("2006-01-02 15:04:05") + ": " + url + " is down!")
}
