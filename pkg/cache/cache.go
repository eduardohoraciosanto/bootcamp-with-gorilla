package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-redis/redis/v8"
)

type Cache interface {
	Set(key string, value interface{}) error
	Get(key string, here interface{}) error
	Del(key string) error
	Alive() bool
}

type redisCache struct {
	client *redis.Client
	ttl    time.Duration
	logger log.Logger
}

func NewRedisCache(logger log.Logger, ttl time.Duration, client *redis.Client) Cache {
	return &redisCache{
		client: client,
		ttl:    ttl,
		logger: log.With(logger, "service", "Cache"),
	}
}

func (c *redisCache) Set(key string, value interface{}) error {
	b, err := json.Marshal(value)
	if err != nil {
		level.Error(c.logger).Log("cache_error", err)
		return err
	}
	level.Info(c.logger).Log("cache_action", "Saving Value to Key", "key", key, "value", value)
	err = c.client.Set(context.Background(), key, string(b), c.ttl).Err()
	if err != nil {
		level.Error(c.logger).Log("cache_error", err)
		return err
	}
	return nil
}

func (c *redisCache) Get(key string, here interface{}) error {
	level.Info(c.logger).Log("cache_action", "Retrieving Key", "key", key)
	val, err := c.client.Get(context.Background(), key).Result()
	if err != nil {
		level.Error(c.logger).Log("cache_error", err)
		return err
	}
	err = json.Unmarshal([]byte(val), here)
	if err != nil {
		level.Error(c.logger).Log("cache_error", err)
		return err
	}
	return nil
}

func (c *redisCache) Del(key string) error {
	level.Info(c.logger).Log("cache_action", "Deleting Key", "key", key)
	numErased, err := c.client.Del(context.Background(), key).Result()
	if err != nil {
		level.Error(c.logger).Log("cache_error", err)
		return err
	}
	if numErased == 0 {
		level.Error(c.logger).Log("cache_error", "Key not Found")
		return redis.Nil
	}

	return nil
}

func (c *redisCache) Alive() bool {
	level.Info(c.logger).Log("cache_action", "Pinging Server")
	if c.client.Ping(context.Background()).Err() != nil {
		level.Error(c.logger).Log("cache_error", "Cache not connected")
		return false
	}
	return true
}
