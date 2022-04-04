package cachedb

import (
	"errors"
	"sync"
)

var ErrorItemNotFound = errors.New("item not found")

type CacheDb struct {
	sync.Mutex
	items map[string]interface{}
}

func NewCache() *CacheDb {
	return &CacheDb{
		items: make(map[string]interface{}),
	}
}

func (c *CacheDb) Set(key string, value interface{}) {
	c.Lock()

	defer c.Unlock()

	c.items[key] = value

}

func (c *CacheDb) Get(key string) interface{} {
	c.Lock()

	defer c.Unlock()

	item, found := c.items[key]

	if !found {
		return nil
	}

	return item
}

func (c *CacheDb) Delete(key string) error {
	c.Lock()

	defer c.Unlock()

	_, found := c.items[key]

	if !found {
		return ErrorItemNotFound
	}

	delete(c.items, key)

	return nil
}
