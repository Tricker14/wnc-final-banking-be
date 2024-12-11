package beanimplement

import (
	"context"
	"time"

	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/bean"
	"github.com/go-redis/redis/v8"
)

type RedisService struct {
	client *redis.Client
}

func NewRedisService() bean.RedisCLient {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return &RedisService{client: client}
}

func (r *RedisService) Set(ctx context.Context, key string, value interface{}, ttl int64) error {
	duration := time.Duration(ttl) * time.Second
	return r.client.Set(ctx, key, value, duration).Err()
}

func (r *RedisService) Get(ctx context.Context, key string) (string, error) {
	value, err := r.client.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}	
	return value, nil
}

func (r *RedisService) Delete(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
