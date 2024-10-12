package services

import (
	"R-I-S-H-A-B-H-S-I-N-G-H/go-microservice/utils/mail_util"
	"fmt"
	"os"
	"strings"
)

type MailService struct {
}

func SendMailService(from string, to string, subject string, body string) error {
	pixel_url := fmt.Sprintf("%s/pixel/123", os.Getenv("SERVICE_BASE_URL"))
	script_tag := fmt.Sprintf("<script src=\"%s\"></script>", pixel_url)
	body = strings.ReplaceAll(body, "{SCRIPT_TAG}", script_tag)
	return mail_util.SendMail(from, to, subject, body)
}
