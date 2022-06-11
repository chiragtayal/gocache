package lru

import (
	"container/list"
	"testing"
	"time"

	"github.com/chiragtayal/gocache/cache"
	"github.com/stretchr/testify/assert"
)

func mockLRU(size int) *LRU {
	return &LRU{
		capacity:     size,
		evictionList: list.New(),
		items:        make(map[string]*list.Element),
	}
}

func TestLRU(t *testing.T) {

	t.Run("Without TTL", func(t *testing.T) {
		lru := mockLRU(2)
		lru.Put("a", "1")
		v, err := lru.Get("a")
		assert.NoError(t, err)
		assert.Equal(t, "1", v)
	})

	t.Run("Cache Entry not present", func(t *testing.T) {
		lru := mockLRU(2)
		v, err := lru.Get("c")
		assert.Error(t, err)
		assert.Empty(t, v, "cache entry should not be present")
	})

	t.Run("Cache entry more then capacity", func(t *testing.T) {
		lru := mockLRU(2)
		lru.Put("b", "2")
		lru.Put("c", "3")
		lru.Put("d", "4")
		//fmt.Println(lru.items)
		v, err := lru.Get("b")
		assert.Error(t, err)
		assert.Empty(t, v, "cache entry should not be present")
	})

	t.Run("With TTL", func(t *testing.T) {
		lru := mockLRU(2)
		lru.Put("d", "4", cache.WithTTL(time.Duration(1*time.Second)))
		time.Sleep(2 * time.Second)
		v, err := lru.Get("d")
		assert.Error(t, err)
		assert.Empty(t, v, "value should be empty")
	})
}
