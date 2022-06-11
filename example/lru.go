package main

import (
	"fmt"
	cache2 "github.com/chiragtayal/gocache/cache"
	_ "github.com/chiragtayal/gocache/cache/lru"
)

func main() {

	cfg := cache2.Config{
		Policy: cache2.LRU,
		Size:   10,
	}

	cache := cache2.NewCache(cfg)

	cache.Put("a", "1")
	cache.Put("b", "2")
	cache.Put("c", "3")
	cache.Put("d", "4")
	cache.Put("e", "5")
	cache.Put("f", "6")
	cache.Put("a", "7")
	cache.Put("h", "8")
	cache.Put("i", "9")
	cache.Put("j", "10")
	cache.Put("k", "11")
	cache.Put("a", "12")

	fmt.Println(cache.Get("i"))
	fmt.Println(cache.Get("k"))
	fmt.Println(cache.Get("c"))
}
