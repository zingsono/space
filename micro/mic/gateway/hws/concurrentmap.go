package hws

import (
	"hash/fnv"
	"sync"
)

// 使用Map数组，实现并发访问Map，数组长度代表分片数量
type ConcurrentMap []*MapShared

type MapShared struct {
	items map[string]interface{}
	sync.RWMutex
}

func (c ConcurrentMap) shard(k string) *MapShared {
	hash32 := fnv.New32()
	hash32.Write([]byte(k))
	i := int32(hash32.Sum32()) % int32(len(c))
	shared := c[i]
	if shared.items == nil {
		c[i] = &MapShared{items: make(map[string]interface{})}
	}
	return shared
}

func (c ConcurrentMap) Put(k string, v interface{}) {
	shared := c.shard(k)
	shared.Lock()
	shared.items[k] = v
	shared.Unlock()
}

func (c ConcurrentMap) Get(k string) interface{} {
	return c.shard(k).items[k]
}

func (c ConcurrentMap) Del(k string) {
	shared := c.shard(k)
	shared.Lock()
	delete(shared.items, k)
	shared.Unlock()
}

func (c ConcurrentMap) Clean() {
	for i := 0; i < len(c); i++ {
		c[i].items = nil
	}
}
