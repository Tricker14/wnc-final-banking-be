package beanimplement

import (
	"context"
	"strconv"

	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/bean"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/utils/env"
	"gopkg.in/gomail.v2"
)

type MailService struct {
	host     string
	port     int
	username string
	password string
	from     string
}

func NewMailService() bean.MailCLient {
	host, _ := env.GetEnv("MAIL_HOST")
	portStr, _ := env.GetEnv("MAIL_PORT")
	port, _ := strconv.Atoi(portStr)
	username, _ := env.GetEnv("MAIL_USERNAME")
	password, _ := env.GetEnv("MAIL_PASSWORD")
	from, _ := env.GetEnv("MAIL_FROM")

	return &MailService{
		host:     host,
		port:     port,
		username: username,
		password: password,
		from:     from,
	}
}

func (m *MailService) SendEmail(ctx context.Context, to string, subject string, body string) error {
	message := gomail.NewMessage()
    message.SetHeader("From", m.from)
    message.SetHeader("To", to)
    message.SetHeader("Subject", subject)
    message.SetBody("text/html", body)

    dialer := gomail.NewDialer(m.host, m.port, m.username, m.password)
    return dialer.DialAndSend(message)
}