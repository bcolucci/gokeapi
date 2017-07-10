package helpers

import "github.com/DannyBen/filecache"

type CacheWrapper struct {
	instance *filecache.Handler
}

func NewCache(dir string, ttlMinutes float64) *CacheWrapper {
	instance := &filecache.Handler{Dir: dir, Life: ttlMinutes}
	return &CacheWrapper{instance}
}

func (c *CacheWrapper) Get(key string) []byte {
	return c.instance.Get(key)
}

func (c *CacheWrapper) Set(key string, bytes []byte) {
	c.instance.Set(key, bytes)
}
