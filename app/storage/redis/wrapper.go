package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type ConnectionWrapper interface {
	redis.Scripter
	Ping(ctx context.Context) *redis.StatusCmd
	HGetAll(ctx context.Context, key string) *redis.MapStringStringCmd
	Del(ctx context.Context, key ...string) *redis.IntCmd
	HTTL(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd
	HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd
	HExpire(ctx context.Context, key string, expiration time.Duration, fields ...string) *redis.IntSliceCmd
	HKeys(ctx context.Context, key string) *redis.StringSliceCmd
	FlushDB(ctx context.Context) *redis.StatusCmd
	Close() error
}

func newStandaloneWrapper(client *redis.Client) ConnectionWrapper {
	return &connectionStandaloneWrapper{client: client}
}

type connectionStandaloneWrapper struct {
	client *redis.Client
}

func (c *connectionStandaloneWrapper) Ping(ctx context.Context) *redis.StatusCmd {
	return c.client.Ping(ctx)
}

func (c *connectionStandaloneWrapper) Close() error {
	return c.client.Close()
}

func (c *connectionStandaloneWrapper) FlushDB(ctx context.Context) *redis.StatusCmd {
	return c.client.FlushDB(ctx)
}

func (c *connectionStandaloneWrapper) HExpire(ctx context.Context, key string, expiration time.Duration, fields ...string) *redis.IntSliceCmd {
	return c.client.HExpire(ctx, key, expiration, fields...)
}

func (c *connectionStandaloneWrapper) HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	return c.client.HSet(ctx, key, values...)
}

func (c *connectionStandaloneWrapper) HTTL(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	return c.client.HTTL(ctx, key, fields...)
}

func (c *connectionStandaloneWrapper) HKeys(ctx context.Context, key string) *redis.StringSliceCmd {
	return c.client.HKeys(ctx, key)
}

func (c *connectionStandaloneWrapper) Del(ctx context.Context, key ...string) *redis.IntCmd {
	return c.client.Del(ctx, key...)
}

func (c *connectionStandaloneWrapper) HGetAll(ctx context.Context, key string) *redis.MapStringStringCmd {
	return c.client.HGetAll(ctx, key)
}

func (c *connectionStandaloneWrapper) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	return c.client.Eval(ctx, script, keys, args...)
}

func (c *connectionStandaloneWrapper) EvalRO(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	return c.client.EvalRO(ctx, script, keys, args...)
}

func (c *connectionStandaloneWrapper) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	return c.client.EvalSha(ctx, sha1, keys, args...)
}

func (c *connectionStandaloneWrapper) EvalShaRO(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	return c.client.EvalShaRO(ctx, sha1, keys, args...)
}

func (c *connectionStandaloneWrapper) ScriptExists(ctx context.Context, hashes ...string) *redis.BoolSliceCmd {
	return c.client.ScriptExists(ctx, hashes...)
}

func (c *connectionStandaloneWrapper) ScriptLoad(ctx context.Context, script string) *redis.StringCmd {
	return c.ScriptLoad(ctx, script)
}

func newClusterWrapper(client *redis.ClusterClient) ConnectionWrapper {
	return &connectionClusterWrapper{client: client}
}

type connectionClusterWrapper struct {
	client *redis.ClusterClient
}

func (c *connectionClusterWrapper) Ping(ctx context.Context) *redis.StatusCmd {
	return c.client.Ping(ctx)
}

func (c *connectionClusterWrapper) Close() error {
	return c.client.Close()
}

func (c *connectionClusterWrapper) FlushDB(ctx context.Context) *redis.StatusCmd {
	return c.client.FlushDB(ctx)
}

func (c *connectionClusterWrapper) HExpire(ctx context.Context, key string, expiration time.Duration, fields ...string) *redis.IntSliceCmd {
	return c.client.HExpire(ctx, key, expiration, fields...)
}

func (c *connectionClusterWrapper) HGetAll(ctx context.Context, key string) *redis.MapStringStringCmd {
	return c.client.HGetAll(ctx, key)
}

func (c *connectionClusterWrapper) HKeys(ctx context.Context, key string) *redis.StringSliceCmd {
	return c.client.HKeys(ctx, key)
}

func (c *connectionClusterWrapper) Del(ctx context.Context, key ...string) *redis.IntCmd {
	return c.client.Del(ctx, key...)
}

func (c *connectionClusterWrapper) HTTL(ctx context.Context, key string, fields ...string) *redis.IntSliceCmd {
	return c.client.HTTL(ctx, key, fields...)
}

func (c *connectionClusterWrapper) HSet(ctx context.Context, key string, values ...interface{}) *redis.IntCmd {
	return c.client.HSet(ctx, key, values...)
}

func (c *connectionClusterWrapper) Eval(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	return c.client.Eval(ctx, script, keys, args...)
}

func (c *connectionClusterWrapper) EvalRO(ctx context.Context, script string, keys []string, args ...interface{}) *redis.Cmd {
	return c.client.EvalRO(ctx, script, keys, args...)
}

func (c *connectionClusterWrapper) EvalSha(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	return c.client.EvalSha(ctx, sha1, keys, args...)
}

func (c *connectionClusterWrapper) EvalShaRO(ctx context.Context, sha1 string, keys []string, args ...interface{}) *redis.Cmd {
	return c.client.EvalShaRO(ctx, sha1, keys, args...)
}

func (c *connectionClusterWrapper) ScriptExists(ctx context.Context, hashes ...string) *redis.BoolSliceCmd {
	return c.client.ScriptExists(ctx, hashes...)
}

func (c *connectionClusterWrapper) ScriptLoad(ctx context.Context, script string) *redis.StringCmd {
	return c.ScriptLoad(ctx, script)
}
