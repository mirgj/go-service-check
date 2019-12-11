# go-service-check

Small utility to check wether a given endpoint is up or down. Supports multiple notification channels

- console
- SMTP
- SendGrid

## Configuration

Based on the notification channel that you want to use to send notifications you're required to set the following ENV variables to make it work:

### Console

No configuration required

### SMTP

The following ENV are required:

```
SMTP_SERVER=<hostname of the SMTP server>
SMTP_USER=<SMTP username>
SMTP_PASSWORD=<SMTP user's password>
SMTP_PORT=25 <or any other in case>
```

### SendGrid

The following ENV are required:

```
SENDGRID_API_KEY=<send grid API Key (secret)>
```

### Service to be checked

The utility uses a `db.json` which needs to have the following structure to:

```json
{
  "sites": [
    {
      "url": "https://google.com",
      "name": "Google",
      "delay": 5,
      "notifier": "console"
    }
  ]
}
```

The array contains a structure with the following fields:

- `url`: URL and/or API endpoint to check
- `name`: Service name to be used as alias for the given URL
- `delay`: Defines the seconds between one ping and another
- `notifier`: Define the type of notifier to be used (allowed values: `console`, `sendgrid`, `smtp`)

In case the **notifier** is `sendgrid` or `smtp` the following options are required on the specific site:

```json
{
  "options": {
    "from": "from@email.com",
    "to": "to@email.com",
    "body": "{{name}} ({{url}}) was not reachable the last {{delay}} seconds!",
    "subject": "{{name}} is down!"
  }
}
```

The `body` and `subject` can contain the following placeholder that will be replaced with the actual value:

- {{url}}: replaced with the service URL
- {{delay}}: replaced with the configured delay
- {{name}}: replaced with the service name

### Full configuration file

Example of a configuration file with all the given options:

```json
{
  "sites": [
    {
      "url": "https://google.com",
      "name": "Google",
      "delay": 5,
      "notifier": "console"
    },
    {
      "url": "https://facebook.com",
      "name": "Facebook",
      "delay": 5,
      "notifier": "sendgrid",
      "options": {
        "from": "from@email.com",
        "to": "to@email.com",
        "body": "{{name}} ({{url}}) was not reachable the last {{delay}} seconds!",
        "subject": "{{name}} is down!"
      }
    },
    {
      "url": "https://stackoverflow.com",
      "name": "Stackoverflow",
      "delay": 5,
      "notifier": "smtp",
      "options": {
        "from": "from@email.com",
        "to": "to@email.com",
        "body": "{{name}} ({{url}}) was not reachable the last {{delay}} seconds!",
        "subject": "{{name}} is down!"
      }
    }
  ]
}
```
