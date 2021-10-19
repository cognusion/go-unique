package unique

import "sync"

// Unique is a simple, goro-safe deduplicating cache
type Unique struct {
	sync.RWMutex
	cache map[interface{}]bool
}

// New returns a reference to an initialzed Unique
func New() *Unique {
	return &Unique{
		cache: make(map[interface{}]bool),
	}
}

// IsUnique returns true if thing does not exist in the cache, and adds it
func (u *Unique) IsUnique(thing interface{}) bool {
	u.Lock()
	defer u.Unlock()

	if _, ok := u.cache[thing]; ok {
		// exists
		return false
	}
	// doesn't exist
	u.cache[thing] = true
	return true
}

// Things returns an array of the cached, unique things
func (u *Unique) Things() []interface{} {
	u.RLock()
	defer u.RUnlock()

	keys := make([]interface{}, len(u.cache))
	c := 0
	for key := range u.cache {
		keys[c] = key
		c++
	}
	return keys
}
