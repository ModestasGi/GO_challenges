package main

import (
	"fmt"
	"net/smtp"
)

func main() {

	// Sender data.
	from := "myemailaddress@gmail.com"
	password := "emailpassword"

	// Receiver email address.
	to := []string{
		"emailaddresswheretosend@gmail.com",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.

	otherMessage := "Greetings via email from GO"
	message := []byte(otherMessage)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
