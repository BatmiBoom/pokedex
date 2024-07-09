package cache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string]entry
	mux   *sync.RWMutex
}

type entry struct {
	created_at time.Time
	val        []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]entry),
		mux:   &sync.RWMutex{},
	}

	go c.reapLoop(interval)

	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = entry{
		created_at: time.Now().UTC(),
		val:        val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	defer c.mux.RUnlock()

	val, ok := c.cache[key]

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
	for k, v := range c.cache {
		if v.created_at.Before(now.Add(-last)) {
			delete(c.cache, k)
		}
	}
}
