package cache

import (
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"
)

type Cache interface {
	Get(key string) (string, error)
	Set(key string, value interface{}) error
}

type cache struct {
	rc *redis.Client
}

func NewCache(rc *redis.Client) *cache {
	return &cache{rc: rc}
}

func (c *cache) Get(key string) (string, error) {
	return c.rc.Get(context.Background(), key).Result()
}
func (c *cache) Set(key string, value interface{}) error {
	valueBytes, _ := json.Marshal(value)
	c.rc.Set(context.Background(), key, valueBytes, 0)
	return nil
}
