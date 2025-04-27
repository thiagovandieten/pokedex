package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	Entries  map[string]CacheEntry
	Mu       sync.Mutex
	interval time.Duration
}

type CacheEntry struct {
	CreatedAt time.Time
	Val       []byte
}

func NewCache(interval time.Duration) *Cache {
	c := &Cache{
		Entries:  make(map[string]CacheEntry),
		interval: interval,
	}
	go c.reapLoop()
	return c
}

func (c *Cache) AddCache(key string, val []byte) {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	entry := CacheEntry{
		CreatedAt: time.Now(),
		Val:       val,
	}
	c.Entries[key] = entry
}

func (c *Cache) GetCache(key string) ([]byte, bool) {
	entry, ok := c.Entries[key]
	if !ok {
		return []byte{}, false
	}
	return entry.Val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()
	for range ticker.C {
		c.removeStaleEntries()
	}

}

func (c *Cache) removeStaleEntries() {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	now := time.Now()
	for k, v := range c.Entries {
		if now.Sub(v.CreatedAt) > c.interval {
			delete(c.Entries, k)
		}
	}

}
