package beanimplement

import (
	"context"
	"strconv"

	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/bean"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/utils/env"
	"gopkg.in/gomail.v2"
)

type MailCLient struct {
	from     string
	dialer	 *gomail.Dialer
}

func NewMailClient() bean.MailCLient {
	host, _ := env.GetEnv("MAIL_HOST")
	portStr, _ := env.GetEnv("MAIL_PORT")
	port, _ := strconv.Atoi(portStr)
	username, _ := env.GetEnv("MAIL_USERNAME")
	password, _ := env.GetEnv("MAIL_PASSWORD")
	from, _ := env.GetEnv("MAIL_FROM")

	dialer := gomail.NewDialer(host, port, username, password)

	return &MailCLient{
		from:     from,
		dialer:   dialer,
	}
}

func (m *MailCLient) SendEmail(ctx context.Context, to string, subject string, body string) error {
	message := gomail.NewMessage()
    message.SetHeader("From", m.from)
    message.SetHeader("To", to)
    message.SetHeader("Subject", subject)
    message.SetBody("text/html", body)

    return m.dialer.DialAndSend(message)
}