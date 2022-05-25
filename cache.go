package cache

import (
	"time"
)

type Cache struct {
	datas []*OneDate
}

type OneDate struct {
	key   string
	val   string
	check bool
	exist bool
	date  time.Time
}

func NewCache() *Cache {
	return &Cache{}
}

func (c *Cache) Get(key string) (string, bool) {
	startTime := time.Now()

	for _, v := range c.datas {
		if v.key == key {
			if !v.check {
				return v.val, true
			}
			res := startTime.Before(v.date)
			if res {
				return v.val, true
			} else {
				v.exist = false
				return "", false
			}
		}
	}
	return "", false

}

func (c *Cache) Put(key, value string) {
	for _, v := range c.datas {
		if v.key == key {
			v.val = value
			return
		}
	}
	data := &OneDate{
		key:   key,
		val:   value,
		check: false,
		exist: true,
	}
	c.datas = append(c.datas, data)
}

func (c *Cache) Keys() []string {
	var arr []string
	for _, v := range c.datas {
		if v.exist {
			arr = append(arr, v.key)
		}

	}
	return arr
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	for _, v := range c.datas {
		if v.key == key {
			v.val = value
			return
		}
	}
	data := &OneDate{
		key:   key,
		val:   value,
		check: true,
		exist: true,
		date:  deadline,
	}
	c.datas = append(c.datas, data)

}
