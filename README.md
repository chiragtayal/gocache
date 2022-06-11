## GoCache

[![Go Report Card](https://goreportcard.com/badge/github.com/chiragtayal/gocache)](https://goreportcard.com/report/github.com/chiragtayal/gocache)

This library provides a `Cache` interface which can be implemented to provide 
Caches based on various kind of eviction policies

The `lru` package implements a fixed capacity and thread safe Cache which uses `LRU`
(Least Recently Used) as an eviction policy with `Time To Live (TTL)` entry option

### Example
```go
import cache "github.com/chiragtayal/gocache/cache"

cfg := cache.Config{Policy: cache.LRU, Size: 10}
c := cache.NewCache(cfg)

### Put cache entry
c.Put("a", "1")
OR
cache.Put("b", "1", cache2.WithTTL(time.Duration(1 * time.Minute)))

### Get cache entry
value, err := c.Get("a")
```

### Generate new mocks
```
1. go install github.com/golang/mock/mockgen@v1.6.0

2. export PATH=$PATH:$(go env GOPATH)/bin

3. go generate -x ./...
```