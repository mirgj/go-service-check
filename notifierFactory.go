package main

// Notifier interface
type Notifier interface {
	Notify(service Service)
}

var notifiers = map[string]Notifier{
	"smtp":     &SMTP{},
	"sendgrid": &SendGrid{},
	"console":  &Console{},
}

func getNotifier(nType string) Notifier {
	if notifer, ok := notifiers[nType]; ok {
		return notifer
	}

	return notifiers["console"]
}
