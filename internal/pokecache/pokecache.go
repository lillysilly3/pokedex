package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entry	map[string]cacheEntry
	mu			sync.Mutex
}

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

func NewCache(interval time.Duration) *Cache{
	cacheMap := make(map[string]cacheEntry)
	c := &Cache{
		entry: cacheMap,
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	c.entry[key] = cacheEntry{
		createdAt:	time.Now(),
		val:		val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, ok := c.entry[key]
	return value.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	
	loopTime := time.NewTicker(interval)
	for range loopTime.C {
		c.mu.Lock()
		for key, value := range c.entry {
			if time.Since(value.createdAt) > interval {
				delete(c.entry, key)
			}
		}
		c.mu.Unlock()
	}
}
