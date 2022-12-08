package config

import (
	"fmt"
	mail "github.com/xhit/go-simple-mail/v2"
	"log"
	"os"
)

var SmtpClient *mail.SMTPClient

func SmtpConfig() {
	server := mail.NewSMTPClient()
	server.Host = "smtp.gmail.com"
	server.Port = 587
	server.Username = os.Getenv("SMTP_EMAIL")
	server.Password = os.Getenv("SMTP_PASSWORD")
	server.Encryption = mail.EncryptionTLS

	smtpClient, err := server.Connect()
	if err != nil {
		fmt.Println("inside error")
		log.Fatal(err)
	}
	SmtpClient = smtpClient
	fmt.Println(SmtpClient)
}
