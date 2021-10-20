// package unique provides a simple, goro-safe deduplicating cache,
// using an RWMutex around a map[interface{}]bool.
// While sync.Map is a simpler solution, it fails to perform
// competitively unless the liklihood of duplicates is very low.
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

/*
// UniqueMap is a simple, goro-safe deduplicating cache.
// It will only be faster if dupes are very rare.
type UniqueMap struct {
	um sync.Map
}

// NewMap returns a reference to an initialzed UniqueMap
func NewMap() *UniqueMap {
	return &UniqueMap{}
}

// IsUnique returns true if thing does not exist in the cache, and adds it
func (u *UniqueMap) IsUnique(thing interface{}) bool {
	_, loaded := u.um.LoadOrStore(thing, true)
	return loaded
}

// Things returns an array of the cached, unique things
func (u *UniqueMap) Things() []interface{} {
	t := make([]interface{}, 0)

	mf := func(key, value interface{}) bool {
		t = append(t, key)
		return true
	}
	u.um.Range(mf)

	return t
}
*/
