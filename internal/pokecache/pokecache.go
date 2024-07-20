// pokecache/cache.go

package pokecache

import (
	"time"
)

type Cache struct {
	cache    map[string]cacheEntry
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache {
		cache: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		val:	val,
		createdAt: time.Now().UTC(),
	}
}


func (c *Cache) Get(key string) ([]byte, bool) {
	cacheE, ok := c.cache[key]
	return cacheE.val, ok
}


func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(interval)
	}
}


func (c *Cache) reap(interval time.Duration) {
	timePassed := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.createdAt.Before(timePassed) {
			delete(c.cache, k)
		}
	}
}