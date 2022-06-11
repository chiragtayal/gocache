## GoCache

This library provides a `Cache` interface which can be implemented to provide 
Caches based on various kind of eviction policies

The `lru` package implements a fixed capacity and thread safe Cache which uses `LRU`
(Least Recently Used) as an eviction policy

Currently, key and value type is string which can be changed to use `interface{}` to 
support `ANY` kind for key and value

### Example
```go

import cache "github.com/chiragtayal/gocache/cache"

cfg := cache.Config{Policy: cache.LRU, Size: 10}
c := cache.NewCache(cfg)

c.Put("a", "1")
value, err := c.Get("a")
```

