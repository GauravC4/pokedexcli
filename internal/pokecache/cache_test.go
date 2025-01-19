package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestNewInMemoryCache(t *testing.T) {
	duration := time.Second * 5
	cache := NewInMemoryCache(duration)
	if cache.duration != duration {
		t.Errorf("invalid duration allocated, expected %v got %v", duration, cache.duration)
	}
	if len(cache.store) > 0 {
		t.Errorf("cache already filled !")
	}
}

func TestAddGetInMemory(t *testing.T) {
	const interval = 5 * time.Second
	cache := NewInMemoryCache(interval)
	testingAddAndGet(t, cache)
}

func TestAddGetRedis(t *testing.T) {
	const interval = 5 * time.Second
	cache := NewRedisCache(interval)
	testingAddAndGet(t, cache)
}

func testingAddAndGet(t *testing.T, cache Cache) {
	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewInMemoryCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
