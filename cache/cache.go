package cache

//go:generate mockgen -destination=cache_mock.go -package=cache -source=cache.go

// Cache interface is implemented by all caches
type Cache interface {

	// Put inserts the provided key/value pair into the gocache,
	// making it available for future get() calls
	Put(key, value string)

	// Get returns a value previously provided via Put(). An empty Option
	// may be returned if the requested data was never inserted or is
	// no longer available.
	Get(key string) (string, error)
}

type Initializer interface {
	// Initialize instantiates a particular type of gocache based on eviction policy
	Initialize(size int)
}

type Composer interface {
	Cache
	Initializer
}

// EvictionPolicy defines gocache type (LRU, LFU etc)
type EvictionPolicy string

const (
	LRU EvictionPolicy = "LRU"
)

func (e EvictionPolicy) String() string {
	switch e {
	case LRU:
		return "LRU"
	}
	return ""
}

// Config is cache configuration
type Config struct {
	Policy EvictionPolicy `json:"evictionPolicy"`
	Size   int            `json:"capacity"`
}

var caches map[EvictionPolicy]Composer

func init() {
	caches = make(map[EvictionPolicy]Composer, 0)
}

// RegisterCache registers various implementation of Cache interface based on eviction policy
func RegisterCache(policy EvictionPolicy, cache Composer) {
	caches[policy] = cache
}

// NewCache instantiate gocache based on the configuration
func NewCache(cfg Config) Cache {
	c, ok := caches[cfg.Policy]
	if !ok {
		return nil
	}
	c.Initialize(cfg.Size)
	return c
}
