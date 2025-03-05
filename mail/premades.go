package mail

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func (es *EmailSender) SendErrorMail(receiver string, error error) {
	// Send an email
	err := es.SendEmail(receiver, "RAIDMON Error Occured!", fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>RAIDMON Error Occured</title>
		</head>
		<body>
			<p>Hello RAIDMON User,</p>
			<p>First of all, don't panic! This is just a notification that an error occured in the RAIDMON application. Your RAID sets are probably still fine, but you might want to check on them. However, what is not fine is RAIDMON. We encountered the following error:</p>
			<pre>%v</pre>
			<p>Please check the logs for more information. If this issue persists, please contact us on GitHub.</p>
			<p>Best regards,</p>
			<p>Your RAIDMON Team</p>
		</body>
		</html>`, error))
	if err != nil {
		logrus.Errorf("Error sending email: %v", err)
	}
}
