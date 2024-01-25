package cache

import (
	"errors"
	"log"
	"sync"
	"time"
)

const invalidationTTL = time.Hour * 3

var errNotFound = errors.New("value not found")

type Cache struct {
	m  map[string]string
	mu sync.RWMutex
}

func NewCache() *Cache {
	c := &Cache{m: make(map[string]string)}
	go c.invalidate()

	return c
}

func (c *Cache) invalidate() {
	tt := time.NewTicker(invalidationTTL)

	for {
		select {
		case <-tt.C:
			c.mu.Lock()
			c.m = make(map[string]string)
			c.mu.Unlock()
		}
	}
}

func (c *Cache) Get(key string) (string, error) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, ok := c.m[key]

	if !ok {
		log.Printf("value not found")
		return "", errNotFound
	}

	return value, nil
}

func (c *Cache) Set(key, value string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.m[key] = value
	return nil
}

func (c *Cache) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.m, key)
}
