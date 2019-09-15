package main

// Notifier interface
type Notifier interface {
	Notify(url string)
}

func getNotifier(nType string) Notifier {
	var notifier Notifier = Console{}

	if nType == "smtp" {
		notifier = SMTP{}
	} else if nType == "sendgrid" {
		notifier = SendGrid{}
	}

	return notifier
}
