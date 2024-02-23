package utils

import (
	"articleproject/api/model/dto"
	"articleproject/config"
	"fmt"
	"net/smtp"
)

func SendMail(email dto.UserMail) {
	to := []string{
		email.To,
	}

	msg := []byte("Subject: " + email.Subject + "\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n" +
		email.Body)

	auth := smtp.PlainAuth(
		"",
		config.Config.SMTP.EmailFrom,
		config.Config.SMTP.EmailPassword,
		config.Config.SMTP.Host,
	)
	err := smtp.SendMail(
		config.Config.SMTP.Host+":"+config.Config.SMTP.Port,
		auth,
		config.Config.SMTP.EmailFrom,
		to,
		msg,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Succesfully.")
}
