package pokecache

import (
	"sync"
	"testing"
	"time"
)

func TestNewCache(t *testing.T) {
	cases := []struct {
		time  time.Duration
		isErr bool
	}{
		{time.Duration(-1), true},
		{time.Minute, time.Minute < MinDuration},
		{time.Second * 5, time.Second*5 < MinDuration},
		{time.Nanosecond, time.Nanosecond < MinDuration},
	}

	for i, c := range cases {
		cache, err := NewCache(c.time)

		if (err != nil) != c.isErr {
			t.Errorf("Error expected for given time %v. Testing case: %v", c.time, i)
			continue
		}

		if c.isErr {
			continue
		}

		if MinDuration > c.time {
			t.Errorf(`interval (%v) should not be less than minDuration (%v), 
					  yet error has not been thrown`, c.time, MinDuration)
			continue
		}

		if cache.maxDuration != c.time {
			t.Errorf("maxDuration (%v) doesn't match given interval (%v)", cache.maxDuration, c)
			continue
		}
	}
}

func TestAdd(t *testing.T) {
	cases := []struct {
		key string
		val []byte
	}{
		{"", []byte{}},
		{"yos", []byte{1, 4, 2}},
		{"empty", []byte{}},
	}

	for i, c := range cases {
		cache := Cache{
			Entries:     make(map[string]cacheEntry),
			Mut:         sync.Mutex{},
			maxDuration: time.Minute,
		}

		cache.Add(c.key, c.val)

		if len(cache.Entries) == 0 {
			t.Errorf("No cache entries. Testing case: %v", i)
			continue
		}

		entry, exists := cache.Entries[c.key]
		if exists == false {
			t.Errorf("Entry not found for key '%v'", c.key)
			continue
		}

		if len(entry.val) != len(c.val) {
			t.Errorf("Different lenghts of values arrays (%v vs %v). Testing case: %v",
				len(entry.val), len(c.val), i)
			continue
		}
	}
}

func TestGet(t *testing.T) {
	cases := []struct {
		key   string
		entry cacheEntry
	}{
		{"", cacheEntry{val: []byte{}}},
		{"yos", cacheEntry{val: []byte{1, 4, 2}}},
		{"empty", cacheEntry{val: []byte{}}},
	}

	cache := Cache{
		Entries:     make(map[string]cacheEntry),
		Mut:         sync.Mutex{},
		maxDuration: time.Minute,
	}

	for i, c := range cases {
		cache.Entries[c.key] = c.entry
		entryVal, hasValue := cache.Get(c.key)
		if !hasValue {
			t.Errorf("Value not found. Testing case: %v", i)
			continue
		}

		if len(entryVal) != len(c.entry.val) {
			t.Errorf("Length of values is different. Testing case: %v", i)
			continue
		}
	}
}
