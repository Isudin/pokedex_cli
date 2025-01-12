package pokecache

import (
	"fmt"
	"sync"
	"time"
)

const MinDuration = 30 * time.Second
const tickDuration = 1 * time.Second

type Cache struct {
	Entries     map[string]cacheEntry
	Mut         sync.Mutex
	Inniciated  bool
	maxDuration time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) (*Cache, error) {
	if interval < MinDuration {
		return &Cache{}, fmt.Errorf("interval (%v) cannot be lower than %v", interval, MinDuration)
	}

	cache := Cache{
		Entries:     make(map[string]cacheEntry),
		Mut:         sync.Mutex{},
		maxDuration: interval,
		Inniciated:  true,
	}

	go cache.reapLoop()

	return &cache, nil
}

func (c *Cache) Add(key string, val []byte) {
	fmt.Println("---SAVING TO CACHE---")
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
