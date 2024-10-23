package cache

import (
	"context"
	"sync"
	"time"
)

type Cache struct {
	mu       sync.Mutex
	interval time.Duration
	entry    map[string]cacheEntry
}

func (c *Cache) Add(url string, val []byte) {
	c.mu.Lock()

	c.entry[url] = cacheEntry{
		time.Now(),
		val,
	}

	c.mu.Unlock()
}

func (c *Cache) Get(url string) ([]byte, bool) {
	if cache, ok := c.entry[url]; ok {
		return cache.val, true
	}

	return nil, false
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for {
		select {
		case <-ticker.C:
			c.removeExpired()
		case <-context.Background().Done():
			ticker.Stop()
			return
		}
	}
}

func (c *Cache) removeExpired() {
	now := time.Now()

	c.mu.Lock()
	for key, cache := range c.entry {
		if now.Sub(cache.createdAt) >= c.interval {
			delete(c.entry, key)
		}
	}
	c.mu.Unlock()
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(time time.Duration) *Cache {
	cache := &Cache{
		interval: time,
		entry:    make(map[string]cacheEntry),
	}

	go cache.reapLoop()

	return cache
}
