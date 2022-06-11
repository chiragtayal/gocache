package lru

import (
	"container/list"
	"github.com/stretchr/testify/assert"
	"testing"
)

func mockLRU(size int) *LRU {
	return &LRU{
		capacity:     size,
		evictionList: list.New(),
		items:        make(map[string]*list.Element),
	}
}

func TestLRU(t *testing.T) {
	lru := mockLRU(2)

	lru.Put("a", "1")
	v, err := lru.Get("a")
	assert.NoError(t, err)
	assert.Equal(t, "1", v)

	lru.Put("b", "2")
	v, err = lru.Get("b")
	assert.NoError(t, err)
	assert.Equal(t, "2", v)
	lru.peek()

	_, _ = lru.Get("a")
	lru.peek()

	lru.Put("c", "3")
	v, err = lru.Get("c")
	assert.NoError(t, err)
	assert.Equal(t, "3", v)
	lru.peek()
}
