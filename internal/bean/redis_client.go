package bean

import "context"

type RedisCLient interface {
	Set(ctx context.Context, key string, value interface{}, ttl int64) error
	Get(ctx context.Context, key string) (string, error)
	Delete(ctx context.Context, key string) error
}