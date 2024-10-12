package services

import "R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/mail_util"

type MailService struct {
}

func SendMailService(from string, to string, subject string, body string) error {
	return mail_util.SendMail(from, to, subject, body)
}
