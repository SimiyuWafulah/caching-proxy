package cache

import (
    "sync"
    "time"
)

type CachedResponse struct {
    StatusCode int
    Header     map[string][]string
    Body       []byte
    Timestamp  time.Time
}

type Cache struct {
    data map[string]CachedResponse
    mutex   sync.RWMutex
}

func NewCache() *Cache {
    return &Cache{
        data: make(map[string]CachedResponse),
    }
}