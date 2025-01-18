package pokecache

import (
	"log"
	"time"
)

type Cache struct {
	duration time.Duration
	store    map[string]CacheEntry
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(duration time.Duration) Cache {
	cache := Cache{
		duration: duration,
		store:    make(map[string]CacheEntry, 1),
	}
	cache.reapLoop()
	return cache
}

func (c *Cache) Get(key string) ([]byte, bool) {
	entry, ok := c.store[key]
	if !ok {
		//TODO: create a logger instance
		log.Printf("pokecache Get : %v miss", key)
		return nil, false
	}
	log.Printf("pokecache Get : %v hit", key)
	return entry.val, true
}

func (c *Cache) Add(key string, val []byte) {
	if entry, ok := c.store[key]; ok {
		entry.createdAt = time.Now()
		localTime := entry.createdAt.Local()
		log.Printf("pokecache Add : %v already exists!, created at updated to %v:%v", key, localTime.Hour(), localTime.Second())
		return
	}
	entry := CacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.store[key] = entry
	localTime := entry.createdAt.Local()
	log.Printf("pokecache Add : %v added at %v:%v", key, localTime.Hour(), localTime.Second())
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.duration)
	go func() {
		for t := range ticker.C {
			localTime := t.Local()
			log.Printf("pokecache reapLoop running at %v:%v", localTime.Hour(), localTime.Second())
			for key, entry := range c.store {
				if time.Since(entry.createdAt) > c.duration {
					createdAtLocalTime := entry.createdAt.Local()
					log.Printf("pokecache reapLoop cleaning key %v created at %v:%v", key, createdAtLocalTime.Hour(), createdAtLocalTime.Second())
					delete(c.store, key)
				}
			}
		}
	}()

}
