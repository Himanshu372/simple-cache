package src

import (
	"sort"
)

type LocalCache struct {
	c    map[string]string
	ru   map[string]int
	size int
}

func NewLocalCache(size int) (*LocalCache, error) {
	c := make(map[string]string, size)
	ru := make(map[string]int, size)
	return &LocalCache{
		c:    c,
		ru:   ru,
		size: size,
	}, nil
}

func (lc *LocalCache) Resize() {
	if len(lc.c) == lc.size {
		// Run eviction and evict top 10 items
		lc.evict()
	}
	return
}

func (lc *LocalCache) Set(key string, val string) {
	lc.Resize()
	// Run eviction and store latest get key
	lc.c[key] = val
	lc.ru[key] = 0
	return
}

func (lc *LocalCache) Get(key string) string {
	if val, ok := lc.c[key]; ok {
		lc.ru[key] += 1
		return val
	}
	return ""
}

func (lc *LocalCache) Delete(key string) {
	delete(lc.c, key)
	delete(lc.ru, key)
}

type KeyUsedCount struct {
	key string
	ruc int
}

func (lc *LocalCache) evict() {
	// Eviction is based on LRU
	kuc := make([]*KeyUsedCount, 0)
	for k, count := range lc.ru {
		kuc = append(kuc, &KeyUsedCount{k, count})
	}
	sort.Slice(kuc, func(i, j int) bool {
		return kuc[i].ruc < kuc[j].ruc
	})
	delete(lc.c, kuc[0].key)
	delete(lc.ru, kuc[0].key)
	return
}
