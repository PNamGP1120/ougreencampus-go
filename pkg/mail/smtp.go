package mail

import (
	"fmt"
	"net/smtp"
)

type SMTPConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	From     string
}

func SendMail(cfg SMTPConfig, to, subject, body string) error {
	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	msg := []byte(
		"From: " + cfg.From + "\r\n" +
			"To: " + to + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-version: 1.0;\r\n" +
			"Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n" +
			body,
	)

	auth := smtp.PlainAuth("", cfg.Username, cfg.Password, cfg.Host)
	return smtp.SendMail(addr, auth, cfg.From, []string{to}, msg)
}
