package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/alfisar/jastip-import/helpers/errorhandler"

	"github.com/go-redis/redis/v8"
)

type redisRepository struct{}

func NewRedisRepository() *redisRepository {
	return &redisRepository{}
}

func (r redisRepository) Insert(ctx context.Context, conn *redis.Client, key string, data string, exp time.Duration) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	errData := conn.Set(ctx, key, data, exp).Err()
	if errData != nil {
		err = fmt.Errorf("insert redis error : %w", errData)
		return
	}

	return
}

func (r redisRepository) Get(ctx context.Context, conn *redis.Client, key string) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	data := conn.Get(ctx, key)

	if data.Err() != nil {
		err = fmt.Errorf("get redis error : %w", data.Err())
		return
	}

	result = data.Val()
	return
}

func (r redisRepository) Delete(ctx context.Context, conn *redis.Client, key string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	errData := conn.Del(ctx, key).Err()

	if errData != nil {
		err = fmt.Errorf("insert redis error : %w", errData)
		return
	}

	return
}

func (r redisRepository) Incr(ctx context.Context, conn *redis.Client, key string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	errData := conn.Incr(ctx, key).Err()
	if errData != nil {
		err = fmt.Errorf("incr redis error : %w", errData)
		return
	}

	return
}

func (r redisRepository) Exp(ctx context.Context, conn *redis.Client, key string, exp time.Duration) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	errData := conn.Expire(ctx, key, exp).Err()
	if errData != nil {
		err = fmt.Errorf("exp redis error : %w", errData)
		return
	}

	return
}
