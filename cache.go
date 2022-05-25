package cache

import (
	"sync"
	"time"
)

type Cache struct {
	data sync.Map
}

type OneDate struct {
	val               string
	hasExpireDeadline bool
	deadline          time.Time
}

func NewCache() *Cache {
	return &Cache{
		data: sync.Map{},
	}
}

func (c *Cache) Get(key string) (string, bool) {
	startTime := time.Now()

	val, ok := c.data.Load(key)

	if !ok {
		return "", false
	}

	v := val.(*OneDate)

	if !v.hasExpireDeadline {
		return v.val, true
	}

	res := startTime.Before(v.deadline)

	if res {
		return v.val, true
	}

	c.data.Delete(key)
	return "", false

}

func (c *Cache) Put(key, value string) {
	data := &OneDate{
		val:               value,
		hasExpireDeadline: false,
	}
	c.data.Store(key, data)
}

func (c *Cache) Keys() []string {
	var arr []string
	c.data.Range(func(key, value interface{}) bool {
		arr = append(arr, key.(string))
		return true
	})
	return arr
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	data := &OneDate{
		val:               value,
		hasExpireDeadline: true,
		deadline:          deadline,
	}
	c.data.Store(key, data)
}

// func main() {
// 	cache := NewCache()
// 	cache.Put("one", "1")
// 	fmt.Println(cache.Keys())
// }
