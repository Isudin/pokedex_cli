package cache

import (
	"sync"
	"time"
)

type Cache struct {
	Entry       map[string]cacheEntry
	Mut         sync.Mutex
	maxDuration time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) (cache Cache) {

	return
}

func (c *Cache) Add(key string, val []byte) {

}

func (c *Cache) Get(key string) ([]byte, bool) {

	return nil, false
}
