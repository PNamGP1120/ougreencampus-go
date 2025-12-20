package mail

import (
	"fmt"
	"net/smtp"
)

type SMTPConfig struct {
	Host string
	Port string
	User string
	Pass string
	From string
}

// SendMail sends a plain text email
func SendMail(cfg SMTPConfig, to, subject, body string) error {
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	msg := []byte(
		"From: " + cfg.From + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n\r\n" +
			body,
	)

	auth := smtp.PlainAuth("", cfg.User, cfg.Pass, cfg.Host)
	return smtp.SendMail(addr, auth, cfg.From, []string{to}, msg)
}
