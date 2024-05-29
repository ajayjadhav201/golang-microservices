package cache

import (
	"github.com/go-redis/redis"
)

type Cache interface {
	Get(key string) (string, error)
	Set(key string, value string) error
	Delete(key string) error
	Close() error
}

type cache struct {
	client *redis.Client
}

func NewCache(client *redis.Client) Cache {
	return &cache{client}
}

func (c *cache) Get(key string) (string, error) {
	return c.client.Get(key).Result()
}

func (c *cache) Set(key string, value string) error {
	return c.client.Set(key, value, 0).Err()
}

func (c *cache) Delete(key string) error {
	return c.client.Del(key).Err()
}

func (c *cache) Close() error {
	return c.client.Close()
}
