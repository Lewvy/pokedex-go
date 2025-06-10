package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	Entries map[string]cacheEntry
	mux     *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		Entries: make(map[string]cacheEntry),
		mux:     &sync.Mutex{},
	}

	go c.reapLoop(interval)
	return c
}
func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.Entries[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.Entries[key]
	return val.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.Entries {
		if v.createdAt.Before(now.Add(-last)) {
			delete(c.Entries, k)
		}
	}
}
