package main

import (
	"net/http"
	"time"
)

type chanStructure struct {
	service Service
	isUp    bool
}

func main() {
	s := readFile()
	c := make(chan chanStructure)

	for _, link := range s.Services {
		go checkStatus(link, c)
	}

	for l := range c {
		go func(link chanStructure) {
			if !link.isUp {
				notifier := getNotifier(link.service.Notifier)
				notifier.Notify(link.service.URL)
			}

			time.Sleep(time.Duration(link.service.Delay) * time.Second)
			checkStatus(link.service, c)
		}(l)
	}
}

func checkStatus(link Service, c chan chanStructure) {
	resp, err := http.Get(link.URL)

	if err != nil {
		c <- chanStructure{link, false}
		return
	}

	c <- chanStructure{link, resp.StatusCode == 200}
}
