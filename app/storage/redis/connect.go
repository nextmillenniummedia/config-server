package redis

import (
	"config-server/app/errors"
	"context"
	"time"

	loggergo "github.com/nextmillenniummedia/logger-go"
	"github.com/redis/go-redis/v9"
)

const FROM_REDIS = "redis"

func Connect(
	config RedisConfig,
	prefix string,
	logger loggergo.ILogger,
) (conn *RedisConnection, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	if len(config.Addr) == 0 {
		return nil, errors.RedisConnectionAddressError
	}
	poolSize := 100

	var clientWrapper ConnectionWrapper
	switch config.Mode {
	case CONNECTION_MODE_STANDALONE:
		options := &redis.Options{
			Addr:         config.Addr[0],
			DB:           config.DB,
			PoolSize:     poolSize,
			MinIdleConns: 5,
			MaxRetries:   3,
		}
		if len(config.Password) > 0 && len(config.Username) > 0 {
			options.Username = config.Username
			options.Password = config.Password
		}
		client := redis.NewClient(options)
		clientWrapper = newStandaloneWrapper(client)
	case CONNECTION_MODE_CLUSTER:
		options := &redis.ClusterOptions{
			Addrs:        config.Addr,
			PoolSize:     poolSize,
			MinIdleConns: 5,
			MaxRetries:   3,
		}
		if len(config.Password) > 0 && len(config.Username) > 0 {
			options.Username = config.Username
			options.Password = config.Password
		}
		client := redis.NewClusterClient(options)
		clientWrapper = newClusterWrapper(client)
	}
	if err := clientWrapper.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	logger.Clone().From(FROM_REDIS).Info("Start", "mode", config.Mode)
	conn = &RedisConnection{
		Client: clientWrapper,
		prefix: prefix,
	}
	return conn, nil
}
