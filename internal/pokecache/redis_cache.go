package pokecache

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	duration time.Duration
	client   *redis.Client
	ctx      context.Context
	keyspace string
}

func NewRedisCache(duration time.Duration) *RedisCache {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	return &RedisCache{
		duration: duration,
		client:   rdb,
		ctx:      context.Background(),
		keyspace: "pokecache",
	}
}

func (c *RedisCache) Get(key string) ([]byte, bool) {
	key = c.keyspace + ":" + key
	val, err := c.client.Get(c.ctx, key).Result()
	if err != nil {
		log.Printf("pokecache(redis) Get : %v miss", key)
		return nil, false
	}
	log.Printf("pokecache(redis) Get : %v hit", key)
	return []byte(val), true
}

func (c *RedisCache) Add(key string, val []byte) {
	key = c.keyspace + ":" + key
	err := c.client.Set(c.ctx, key, val, c.duration).Err()
	if err != nil {
		log.Printf("pokecache(redis) Add : error storing key : %v, error : %v", key, err)
	}
	localTime := time.Now().Local()
	log.Printf("pokecache(redis) Add : %v added at %v:%v", key, localTime.Hour(), localTime.Second())
}
