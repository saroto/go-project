package mail

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendMail(recipient, subject, body string) error {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUser := os.Getenv("SMTP_USER")
	smtpPass := os.Getenv("SMTP_PASSWORD")

	auth := smtp.PlainAuth("", smtpUser, smtpPass, smtpHost)

	addr := smtpHost + ":" + smtpPort
	hearder := "From: " + smtpUser + "\r\n" +
		"To: " + recipient + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIMI-Version: 1.0; \r\n" +
		"Content-Type: text/html; charset=UTF-8;\r\n\r\n"
	msg := []byte(hearder + body)
	if err := smtp.SendMail(addr, auth, smtpUser, []string{recipient}, msg); err != nil {
		fmt.Println("Error sending email", err)
		return err
	}
	return nil
}
