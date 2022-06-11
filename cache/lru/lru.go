package lru

import (
	"container/list"
	"fmt"
	"github.com/chiragtayal/gocache/cache"
	"sync"
)

type LRU struct {
	sync.Mutex

	capacity     int
	evictionList *list.List
	items        map[string]*list.Element
}

type entry struct {
	key, value string
}

func init() {
	cache.RegisterCache(cache.LRU, &LRU{})
}

// Initialize instantiates an LRU gocache of the provided capacity
func (c *LRU) Initialize(size int) {
	c.capacity = size
	c.evictionList = list.New()
	c.items = make(map[string]*list.Element)
}

// Put adds an entry to gocache
func (c *LRU) Put(key, value string) {
	c.Lock()
	defer c.Unlock()

	// check if key exists
	if e, ok := c.items[key]; ok {
		c.evictionList.MoveToFront(e)
		e.Value.(*entry).value = value
		return
	}

	// create a new entry and add to gocache
	e := entry{key: key, value: value}
	item := c.evictionList.PushFront(&e)
	c.items[key] = item

	// check if capacity is exceeded, evict lru item
	if c.capacity < c.evictionList.Len() {
		e := c.evictionList.Back()
		c.evictionList.Remove(e)

		kv := e.Value.(*entry)
		delete(c.items, kv.key)
	}

	return
}

// Get looks up a key's value from the gocache
func (c *LRU) Get(key string) (string, error) {
	c.Lock()
	defer c.Unlock()

	// check if key exists
	e, ok := c.items[key]
	if !ok {
		return "", fmt.Errorf("gocache entry not present for key: %s", key)
	}

	// if exists, move to front
	c.evictionList.MoveToFront(e)

	return e.Value.(*entry).value, nil
}

func (c *LRU) peek() {
	e := c.evictionList.Front()
	fmt.Println("key: ", e.Value.(*entry).key, ", Value:", e.Value.(*entry).value)
}
