package main

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// SendGrid notifier
type SendGrid struct {
}

// Notify the user via email
func (e SendGrid) Notify(service Service) {
	t := time.Now()
	fmt.Println(t.Format("2006-01-02 15:04:05") + ": " + service.URL + " is down, sending notification via SendGrid!")

	apiKey := SendGridAPIKey
	sendGridURL := "https://api.sendgrid.com/v3/mail/send"
	var jsonStr = []byte(`
	{
		"personalizations": [
			{
				"to": [
					{
						"email": "` + service.SendGrid.To + `"
					}
				],
				"subject": "` + replaceTokens(service, service.SendGrid.Subject) + `"
			}
		],
		"from": {
			"email": "` + service.SendGrid.From + `"
		},
		"content": [
			{
				"type": "text/plain",
				"value": "` + replaceTokens(service, service.SendGrid.Body) + `"
			}
		]
	}`)

	req, err := http.NewRequest("POST", sendGridURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending the notification: " + err.Error())
	}

	defer resp.Body.Close()
}

func replaceTokens(service Service, str string) string {
	newStr := strings.Replace(str, "{{url}}", service.URL, -1)
	newStr = strings.Replace(newStr, "{{name}}", service.Name, -1)
	newStr = strings.Replace(newStr, "{{delay}}", strconv.Itoa(service.Delay), -1)

	return newStr
}
