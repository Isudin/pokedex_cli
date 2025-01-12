package cache

import (
	"fmt"
	"sync"
	"time"
)

const minDuration = 30 * time.Second

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

	return &cache, nil
}

func (c *Cache) Add(key string, val []byte) {
	c.Entries[key] = cacheEntry{val: val, createdAt: time.Now()}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, isErr := c.Entries[key]
	return entry.val, isErr
}
