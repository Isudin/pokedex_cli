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
	c.Entries[key] = cacheEntry{val: val, createdAt: time.Now()}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, isErr := c.Entries[key]
	return entry.val, isErr
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(tickDuration)
	defer ticker.Stop()
	fmt.Println("Starting loop")

	for {
		fmt.Println("Waiting for first tick")
		time := <-ticker.C
		fmt.Printf("Tick acquired: %v", time)
		for key, entry := range c.Entries {
			timeToClose := entry.createdAt.Add(c.maxDuration)
			if timeToClose.Compare(time) < 1 {
				c.Mut.Lock()
				delete(c.Entries, key)
				c.Mut.Unlock()
			}
		}
		fmt.Println("Resetting ticker")
		ticker.Reset(tickDuration)
	}
}
