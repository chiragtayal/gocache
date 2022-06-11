package lru

import (
	"container/list"
	"fmt"
	"sync"
	"time"

	"github.com/chiragtayal/gocache/cache"
)

type LRU struct {
	sync.Mutex

	capacity     int
	evictionList *list.List
	items        map[string]*list.Element
}

func init() {
	cache.Register(cache.LRU, &LRU{})
}

// Initialize instantiates an LRU cache of the provided capacity
func (c *LRU) Initialize(size int) {
	c.capacity = size
	c.evictionList = list.New()
	c.items = make(map[string]*list.Element)
}

// Put adds an entry to cache
func (c *LRU) Put(key, value string, opts ...cache.Options) {
	c.Lock()
	defer c.Unlock()

	// evict expired cache entry if ttl present
	c.evictExpired()

	// check if key exists
	if e, ok := c.items[key]; ok {
		c.evictionList.MoveToFront(e)
		e.Value.(*cache.Entry).Value = value
		return
	}

	// create a new entry and add to cache
	e := cache.Entry{Key: key, Value: value}
	for _, opt := range opts {
		opt(&e)
	}

	item := c.evictionList.PushFront(&e)
	c.items[key] = item

	// check if capacity is exceeded, evict lru item
	if c.capacity < c.evictionList.Len() {
		e := c.evictionList.Back()
		c.evictionList.Remove(e)

		kv := e.Value.(*cache.Entry)
		delete(c.items, kv.Key.(string))
	}

	return
}

// Get looks up a key's value from the cache
func (c *LRU) Get(key string) (string, error) {
	c.Lock()
	defer c.Unlock()

	// evict expired cache entry if ttl present
	c.evictExpired()

	// check if key exists
	e, ok := c.items[key]
	if !ok {
		return "", fmt.Errorf("cache entry not present for key: %s", key)
	}

	// if exists, move to front
	c.evictionList.MoveToFront(e)

	return e.Value.(*cache.Entry).Value.(string), nil
}

func (c *LRU) peek() {
	e := c.evictionList.Front()
	fmt.Println("key: ", e.Value.(*cache.Entry).Key, ""+
		", Value:", e.Value.(*cache.Entry).Value, " TTL: ", e.Value.(*cache.Entry).TTL)
}

func (c *LRU) evictExpired() {
	for k, v := range c.items {
		e := v.Value.(*cache.Entry)
		if !e.TTL.IsZero() && time.Now().After(e.TTL) {
			fmt.Printf("cache entry expired for key: %v \n", k)
			c.evictionList.Remove(v)
			delete(c.items, k)
		}
	}
}
