package utils

import (
	"container/list"
	"sync"
	"time"
)

type Cache struct {
	m    map[string]*item
	max  int
	list list.List
	sync.RWMutex
}

type item struct {
	value      interface{}
	created    int64
	expiration int64
}

var MemoryCache Cache = Cache{
	m:   make(map[string]*item),
	max: 1000000,
}

func (c *Cache) Get(key string) interface{} {
	c.RLock()
	defer c.RUnlock()

	v, ok := c.m[key]
	if !ok {
		return nil
	}
	if v.expiration > 0 && time.Now().UnixNano() > v.expiration {
		// 缓存项已过期
		c.delete(key)
		return nil
	}
	return v.value
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	c.Lock()
	defer c.Unlock()

	// 检查是否超过最大数量，如果超过则触发清理
	for c.list.Len() >= c.max {
		f := c.list.Front()
		if f != nil {
			c.list.Remove(f)
			delete(c.m, f.Value.(string))
		}
	}

	// 添加新缓存项
	created := time.Now().UnixNano()
	expiration := time.Now().Add(ttl).UnixNano()
	c.m[key] = &item{value, created, expiration}
	c.list.PushBack(key)
}

func (c *Cache) SetMax(max int) {
	c.max = max
	c.cleanup()
}

// 采用FIFO清理策略超出最大数量的缓存项
func (c *Cache) cleanup() {
	for c.list.Len() > c.max {
		f := c.list.Front()
		if f != nil {
			c.Lock()
			c.list.Remove(f)
			delete(c.m, f.Value.(string))
			c.Unlock()
		}
	}
}

func (c *Cache) delete(key string) {
	delete(c.m, key)
	for e := c.list.Front(); e != nil; e = e.Next() {
		if e.Value.(string) == key {
			c.list.Remove(e)
			return
		}
	}
}
