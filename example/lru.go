package main

import (
	"fmt"
	"time"

	cache2 "github.com/chiragtayal/gocache/cache"
	_ "github.com/chiragtayal/gocache/cache/lru"
)

func main() {

	cfg := cache2.Config{
		Policy: cache2.LRU,
		Size:   10,
	}

	cache := cache2.New(cfg)

	cache.Put("a", "1", cache2.WithTTL(time.Duration(1*time.Minute)))
	cache.Put("b", "2")
	cache.Put("c", "3")
	cache.Put("d", "4")
	cache.Put("e", "5", cache2.WithTTL(time.Duration(1*time.Microsecond)))
	cache.Put("f", "6")
	cache.Put("a", "7")
	cache.Put("h", "8")
	cache.Put("i", "9")
	cache.Put("j", "10")
	cache.Put("k", "11")

	fmt.Println(cache.Get("a"))
	fmt.Println(cache.Get("e"))
	fmt.Println(cache.Get("c"))

	time.Sleep(1 * time.Minute)
	fmt.Println(cache.Get("a"))

	fmt.Println(cache.Get("h"))
	fmt.Println(cache.Get("i"))
}
