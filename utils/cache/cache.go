package cache

import "time"

type CacheItem struct {
	Value     string
	ExpiresAt time.Time
}

var cache = make(map[string]CacheItem)

func SetCache(key, value string, ttl time.Duration) {
	cache[key] = CacheItem{Value: value, ExpiresAt: time.Now().Add(ttl)}
}

func GetCache(key string) (string, bool) {
	item, found := cache[key]
	if !found || time.Now().After(item.ExpiresAt) {
		return "", false
	}
	return item.Value, true
}
