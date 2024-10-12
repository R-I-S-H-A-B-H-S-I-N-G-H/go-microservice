package mail_util

import (
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/error_util"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

func SendMail(from string, to string, subject string, body string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	mail_port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	error_util.Handle("Failed to convert string to int", err)

	d := gomail.NewDialer(os.Getenv("MAIL_HOST"), mail_port, os.Getenv("MAIL_USERNAME"), os.Getenv("MAIL_PASSWORD"))
	send_err := d.DialAndSend(m)
	error_util.Handle("Failed to send mail", send_err)
	return err
}
