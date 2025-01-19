package pokecache

import (
	"log"
	"sync"
	"time"
)

type InMemoryCache struct {
	duration time.Duration
	store    map[string]InMemoryCacheEntry
	mux      *sync.RWMutex
}

type InMemoryCacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewInMemoryCache(duration time.Duration) *InMemoryCache {
	cache := InMemoryCache{
		duration: duration,
		store:    make(map[string]InMemoryCacheEntry, 1),
		mux:      &sync.RWMutex{},
	}
	go cache.reapLoop()
	return &cache
}

func (c *InMemoryCache) Get(key string) ([]byte, bool) {
	c.mux.RLock()
	entry, ok := c.store[key]
	c.mux.RUnlock()
	if !ok {
		log.Printf("pokecache(inmemory) Get : %v miss", key)
		return nil, false
	}
	log.Printf("pokecache(inmemory) Get : %v hit", key)
	return entry.val, true
}

func (c *InMemoryCache) Add(key string, val []byte) {
	c.mux.RLock()
	if entry, ok := c.store[key]; ok {
		entry.createdAt = time.Now()
		localTime := entry.createdAt.Local()
		log.Printf("pokecache(inmemory) Add : %v already exists!, created at updated to %v:%v", key, localTime.Hour(), localTime.Second())
		return
	}
	c.mux.RUnlock()
	entry := InMemoryCacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.mux.Lock()
	c.store[key] = entry
	c.mux.Unlock()
	localTime := entry.createdAt.Local()
	log.Printf("pokecache(inmemory) Add : %v added at %v:%v", key, localTime.Hour(), localTime.Second())
}

func (c *InMemoryCache) reapLoop() {
	ticker := time.NewTicker(c.duration)

	for t := range ticker.C {
		localTime := t.Local()
		log.Printf("pokecache(inmemory) reapLoop running at %v:%v", localTime.Hour(), localTime.Second())
		c.mux.Lock()
		for key, entry := range c.store {
			if time.Since(entry.createdAt) > c.duration {
				createdAtLocalTime := entry.createdAt.Local()
				log.Printf("pokecache(inmemory) reapLoop cleaning key %v created at %v:%v", key, createdAtLocalTime.Hour(), createdAtLocalTime.Second())
				delete(c.store, key)
			}
		}
		c.mux.Unlock()
	}

}
