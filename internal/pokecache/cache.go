package cache

import (
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

func NewCache(interval time.Duration) (cache Cache, err error) {

	return
}

func (c *Cache) Add(key string, val []byte) {

}

func (c *Cache) Get(key string) ([]byte, bool) {

	return nil, false
}
