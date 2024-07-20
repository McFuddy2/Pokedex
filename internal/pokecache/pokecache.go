// pokecache/cache.go

package pokecache

import (
	"time"
	"sync"
)

type Cache struct {
	cache    map[string]cacheEntry
	mu 		 *sync.Mutex
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

func NewCache(interval time.Duration) Cache {
	c := Cache {
		cache: 	make(map[string]cacheEntry),
		mu: 	&sync.Mutex{},
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{
		val:	val,
		createdAt: time.Now().UTC(),
	}
}


func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
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
	c.mu.Lock()
	defer c.mu.Unlock()
	timePassed := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.createdAt.Before(timePassed) {
			delete(c.cache, k)
		}
	}
}