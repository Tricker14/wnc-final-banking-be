package beanimplement

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/bean"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/utils/env"
	"gopkg.in/gomail.v2"
)

type MailClient struct {
	from   string
	dialer *gomail.Dialer
}

func NewMailClient() bean.MailClient {
	host, _ := env.GetEnv("MAIL_HOST")
	portStr, _ := env.GetEnv("MAIL_PORT")
	port, _ := strconv.Atoi(portStr)
	username, _ := env.GetEnv("MAIL_USERNAME")
	password, _ := env.GetEnv("MAIL_PASSWORD")
	from, _ := env.GetEnv("MAIL_FROM")

	dialer := gomail.NewDialer(host, port, username, password)

	return &MailClient{
		from:   from,
		dialer: dialer,
	}
}

func (m *MailClient) SendEmail(ctx context.Context, to string, subject string, code string, context string, ttl time.Duration) error {
	message := gomail.NewMessage()

	contextMessage := fmt.Sprintf("Enter the following code to %s:", context)

	body := fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<body>
			<p>You have selected <strong>%s</strong> as your email address.</p>
			<p>%s</p>
			<h2 style="color: #4CAF50;">%s</h2>
			<p>This code will expire %d minute after this email was sent.</p>
			<hr>
			<p><strong>Why you received this email.</strong></p>
			<p>We require verification whenever an email address is updated.</p>
			<p>If you did not make this request, you can ignore this email.</p>
		</body>
		</html>`, to, contextMessage, code, ttl/time.Minute)

	message.SetHeader("From", m.from)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)

	return m.dialer.DialAndSend(message)
}
