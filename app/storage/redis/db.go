package redis

import (
	"context"
)

type RedisConnection struct {
	Client ConnectionWrapper
	prefix string
}

func (repo *RedisConnection) ConfigAdd(ctx context.Context) (err error) {
	return nil
}
