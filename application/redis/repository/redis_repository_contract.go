package repository

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisRepositoryContract interface {
	Insert(ctx context.Context, conn *redis.Client, key string, data string, exp time.Duration) (err error)
	Get(ctx context.Context, conn *redis.Client, key string) (result string, err error)
	Delete(ctx context.Context, conn *redis.Client, key string) (err error)
	Incr(ctx context.Context, conn *redis.Client, key string) (err error)
	Exp(ctx context.Context, conn *redis.Client, key string, exp time.Duration) (err error)
}
