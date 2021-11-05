package cache

import "github.com/go-redis/redis/v8"

type Cache interface {
	Get(key string) (interface{}, error)
	Set(key string, value interface{}) error
}

type cache struct {
	rc *redis.Client
}

func NewCache(rc *redis.Client) *cache {

	return &cache{rc: rc}
}

func (c *cache) Get(key string) (interface{}, error) {
	return nil, nil
}
func (c *cache) Set(key string, value interface{}) error {
	return nil
}
