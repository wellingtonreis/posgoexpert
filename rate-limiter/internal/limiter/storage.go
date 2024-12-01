package limiter

import (
	"context"
	"time"

	redis "github.com/go-redis/redis/v8"
)

type Storage interface {
	Increment(key string) (int, error)
	GetTTL(key string) (int64, error)
	Block(key string, duration int64) error
}

type RedisStorage struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisStorage(addr string) *RedisStorage {
	rdb := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	return &RedisStorage{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (rs *RedisStorage) Increment(key string) (int, error) {
	v, err := rs.client.Incr(rs.ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return int(v), nil
}

func (rs *RedisStorage) GetTTL(key string) (int64, error) {
	ttl, err := rs.client.TTL(rs.ctx, key).Result()
	if err != nil {
		return 0, err
	}
	return int64(ttl.Seconds()), nil
}

func (rs *RedisStorage) Block(key string, duration int64) error {
	return rs.client.Set(rs.ctx, key, 1, time.Duration(duration)*time.Second).Err()
}
