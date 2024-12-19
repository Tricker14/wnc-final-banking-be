package bean

import (
	"context"
	"time"
)

type MailClient interface {
	SendEmail(ctx context.Context, to, subject, code, context string, ttl time.Duration) error
}
