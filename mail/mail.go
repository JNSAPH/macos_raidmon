package mail

import (
	"fmt"
	"net/smtp"

	"github.com/sirupsen/logrus"
)

type EmailSender struct {
	SmtpHost   string
	SmtpPort   int
	Username   string
	Password   string
	Sender     string
	SMTPClient *smtp.Client
}

func NewEmailSender(Host string, Port int, Username string, Password string, Sender string) *EmailSender {
	return &EmailSender{
		SmtpHost:   Host,
		SmtpPort:   Port,
		Username:   Username,
		Password:   Password,
		Sender:     Sender,
		SMTPClient: nil,
	}
}

func (es *EmailSender) SendEmail(recipient, subject, htmlContent string) error {

	auth := smtp.PlainAuth("", es.Username, es.Password, es.SmtpHost)

	message := es.createHTMLMessage(subject, htmlContent)

	err := smtp.SendMail(fmt.Sprintf("%s:%d", es.SmtpHost, es.SmtpPort), auth, es.Sender, []string{recipient}, []byte(message))
	if err != nil {
		return err
	}

	logrus.Info("Email sent to ", recipient)
	return nil
}

func (es *EmailSender) createHTMLMessage(subject, htmlContent string) []byte {
	htmlBody := fmt.Sprintf("<html><body>%s</body></html>", htmlContent)

	message := []byte("Subject: " + subject + "\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\n" +
		"\n" + htmlBody)

	return message
}
