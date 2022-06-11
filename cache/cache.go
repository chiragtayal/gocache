package cache

//go:generate mockgen -destination=cache_mock.go -package=cache -source=cache.go

// Cache interface is implemented by all caches
type Cache interface {

	// Put inserts the provided key/value pair into the cache,
	// making it available for future get() calls
	Put(string, string, ...Options)

	// Get returns a value previously provided via Put(). An empty Option
	// may be returned if the requested data was never inserted or is
	// no longer available.
	Get(string) (string, error)
}

type Initializer interface {
	// Initialize instantiates a particular type of cache based on eviction policy
	Initialize(size int)
}

type Composer interface {
	Cache
	Initializer
}

// EvictionPolicy defines cache type (LRU, LFU etc)
type EvictionPolicy string

const (
	LRU     EvictionPolicy = "LRU"
	LRU_TTL EvictionPolicy = "LRU_TTL"
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

// Register registers various implementation of Cache interface based on eviction policy
func Register(policy EvictionPolicy, cache Composer) {
	caches[policy] = cache
}

// New instantiates Cache based on the configuration
func New(cfg Config) Cache {
	c, ok := caches[cfg.Policy]
	if !ok {
		return nil
	}
	c.Initialize(cfg.Size)
	return c
}
