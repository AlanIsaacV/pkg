package redis

import (
	"context"
	"sync"

	"github.com/gofiber/fiber/v2/log"
	"github.com/redis/go-redis/v9"
)

var (
	client     *redis.Client
	clientOnce sync.Once
)

func Client() *redis.Client {
	clientOnce.Do(
		func() {
			c := Config()
			client = redis.NewClient(&redis.Options{Addr: c.Addr, DB: c.Db})
		},
	)
	return client
}

func Flush(ctx context.Context, pattern string) (int64, error) {
	prefix := Config().Prefix
	rdb := Client()

	if pattern == "" && prefix == "" {
		rdb.FlushDB(ctx)
		return -1, nil
	}
	if pattern == "" {
		pattern = "*"
	}

	pattern = prefix + pattern
	var cursor uint64
	var keysFull []string

	for {
		var keys []string
		var err error

		keys, cursor, err = rdb.Scan(ctx, cursor, pattern, 5000).Result()
		if err != nil {
			log.Error(err)
		}

		keysFull = append(keysFull, keys...)
		if cursor == 0 {
			break
		}
	}
	if len(keysFull) > 0 {
		return rdb.Del(ctx, keysFull...).Result()
	}
	return 0, nil
}
