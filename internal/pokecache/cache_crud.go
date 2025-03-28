package pokecache

import (
	"time"
)

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.cacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	if entry, ok := cache.cacheEntries[key]; ok {
		return entry.val, ok
	}
	return nil, false
}

func (cache *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		<-ticker.C
		cache.mu.Lock()
		for key, entry := range cache.cacheEntries {
			if time.Since(entry.createdAt) > interval {
				delete(cache.cacheEntries, key)
			}
		}
		cache.mu.Unlock()

	}
}
