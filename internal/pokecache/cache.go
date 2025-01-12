package cache

import (
	"fmt"
	"sync"
	"time"
)

const minDuration = 5 * time.Second
const tickDuration = 1 * time.Second

type Cache struct {
	Entries     map[string]cacheEntry
	Mut         sync.Mutex
	maxDuration time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) (*Cache, error) {
	if interval < minDuration {
		return &Cache{}, fmt.Errorf("interval (%v) cannot be lower than %v", interval, minDuration)
	}

	cache := Cache{
		Entries:     make(map[string]cacheEntry),
		Mut:         sync.Mutex{},
		maxDuration: interval,
	}

	go cache.reapLoop()

	return &cache, nil
}

func (c *Cache) Add(key string, val []byte) {
	c.Mut.Lock()
	c.Entries[key] = cacheEntry{val: val, createdAt: time.Now()}
	c.Mut.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.Mut.Lock()
	entry, isErr := c.Entries[key]
	c.Mut.Unlock()
	return entry.val, isErr
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(tickDuration)
	defer ticker.Stop()

	for {
		time := <-ticker.C
		c.Mut.Lock()
		for key, entry := range c.Entries {
			timeToClose := entry.createdAt.Add(c.maxDuration)
			if timeToClose.Compare(time) < 1 {
				delete(c.Entries, key)
			}
		}
		c.Mut.Unlock()
		ticker.Reset(tickDuration)
	}
}
