package pokecache

import (
	"sync"
	"time"
)

type cache struct {
	data map[string]cacheEntry
	mu   *sync.Mutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func newCache(interval time.Duration) cache {
	c := cache{
		data: make(map[string]cacheEntry),
		mu:   &sync.Mutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *cache) add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *cache) get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.data[key]
	return val.val, ok
}

func (c *cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *cache) reap(now time.Time, last time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for key, value := range c.data {
		if value.createdAt.Before(now.Add(-last)) {
			delete(c.data, key)
		}
	}
}
