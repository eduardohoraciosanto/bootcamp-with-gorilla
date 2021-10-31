package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
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
	logger *logrus.Logger
}

func NewRedisCache(logger *logrus.Logger, ttl time.Duration, client *redis.Client) Cache {
	return &redisCache{
		client: client,
		ttl:    ttl,
		logger: logger,
	}
}

func (c *redisCache) Set(key string, value interface{}) error {
	b, err := json.Marshal(value)
	if err != nil {
		c.logger.WithError(err).Error("cache_error")
		return err
	}
	c.logger.WithField("key", key).WithField("value", value).Log(logrus.InfoLevel, "Saving Value to Key")
	err = c.client.Set(context.Background(), key, string(b), c.ttl).Err()
	if err != nil {
		c.logger.WithError(err).Error("cache_error")
		return err
	}
	return nil
}

func (c *redisCache) Get(key string, here interface{}) error {
	c.logger.WithField("key", key).Log(logrus.InfoLevel, "Retrieving Key")
	val, err := c.client.Get(context.Background(), key).Result()
	if err != nil {
		c.logger.WithError(err).Error("cache_error")
		return err
	}
	err = json.Unmarshal([]byte(val), here)
	if err != nil {
		c.logger.WithError(err).Error("cache_error")
		return err
	}
	return nil
}

func (c *redisCache) Del(key string) error {
	c.logger.WithField("key", key).Log(logrus.InfoLevel, "Deleting Key")
	numErased, err := c.client.Del(context.Background(), key).Result()
	if err != nil {
		c.logger.WithError(err).Error("cache_error")
		return err
	}
	if numErased == 0 {
		c.logger.Error("cache key not found")
		return redis.Nil
	}

	return nil
}

func (c *redisCache) Alive() bool {
	c.logger.Log(logrus.InfoLevel, "Pinging Redis")
	if c.client.Ping(context.Background()).Err() != nil {
		c.logger.Error("cache not connected")
		return false
	}
	return true
}
