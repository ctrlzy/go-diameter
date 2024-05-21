package diam

import (
	"reflect"
	"sync"
	"sync/atomic"
)

// static cache for struct fields
var (
	cache = structCache{}
)

// empty here for later improvement
type cachedField struct {
}

type structCache struct {
	lock sync.Mutex
	m    atomic.Value // map[reflect.Type][]cachedField
}

func (sc *structCache) Get(key reflect.Type) ([]cachedField, bool) {
	m := sc.m.Load().(map[reflect.Type][]cachedField)
	c, found := m[key]
	return c, found
}

func (sc *structCache) Set(key reflect.Type, value []cachedField) {
	sc.lock.Lock()
	defer sc.lock.Unlock()

	m := sc.m.Load().(map[reflect.Type][]cachedField)
	nm := make(map[reflect.Type][]cachedField, len(m)+1)
	for k, v := range m {
		nm[k] = v
	}
	nm[key] = value
	sc.m.Store(nm)
}
