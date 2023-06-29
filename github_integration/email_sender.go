package main

import (
	"log"

	"github.com/go-gomail/gomail"
)

func sendEmail(recipient, subject, body string) {
	m := gomail.NewMessage()
	m.SetHeader("From", emailSender)
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", subject)
	m.SetBody("text/plain", body)

	d := gomail.NewDialer(emailSMTPServer, emailSMTPPort, emailSMTPUser, emailSMTPPass)

	err := d.DialAndSend(m)
	if err != nil {
 	 log.Fatalf("Failed to send email: %v", err)
  }
}
