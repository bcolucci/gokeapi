package api

import "github.com/DannyBen/filecache"

type Cache struct {
	intance *filecache.Handler
}

func NewCache(dir string, ttlMinutes float64) *Cache {
	instance := &filecache.Handler{Dir: dir, Life: ttlMinutes}
	return &Cache{instance}
}

func (c *Cache) Get(key string) []byte {
	return c.intance.Get(key)
}

func (c *Cache) Set(key string, bytes []byte) {
	c.intance.Set(key, bytes)
}
