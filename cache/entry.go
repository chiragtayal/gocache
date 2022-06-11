package cache

import (
	"time"
)

const defaultTTL time.Duration = 5 * time.Minute

type Entry struct {
	Key, Value interface{}
	TTL        time.Time
}

type Options func(e *Entry)

func WithTTL(ttl time.Duration) Options {
	return func(e *Entry) {
		if ttl == 0 {
			ttl = defaultTTL
		}
		e.TTL = time.Now().Add(ttl)
	}
}
