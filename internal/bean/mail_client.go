package bean

import "context"

type MailCLient interface {
	SendEmail(ctx context.Context, to, subject, body string) error
}