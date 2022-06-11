// Package cache provides the implementation of eviction policy based Cache
// this package can be extended to implement caches based on other eviction policies
//
// cache in this package take locks while operating, and are therefore thread-safe for consumers
package cache
