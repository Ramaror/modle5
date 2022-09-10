package main

import (
	"fmt"
	"sync"
	"time"
)

type Cache struct {
	storage map[string]int
	mu      sync.Mutex
}

func (c *Cache) Increase(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] += value
}

func (c *Cache) Set(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] = value
}

func (c *Cache) Get(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.storage[key]
}

func (c *Cache) Remove(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	delete(c.storage, key)

}

/*
Перепишите класс Cache так, чтобы он стал thread-safe(потокобезопасным)

package main

import "time"

	type Cache struct {
		storage map[string]int
	}

	func (c *Cache) Increase(key string, value int) {
		c.storage[key] += value
	}

	func (c *Cache) Set(key string, value int) {
		c.storage[key] = value
	}

	func (c *Cache) Get(key string) int {
		return c.storage[key]
	}

	func (c *Cache) Remove(key string) {
		delete(c.storage, key)
	}
*/
func main() {
	bc := &Cache{
		storage: map[string]int{"first": 1},
	}
	go bc.Increase("first", 22)
	go bc.Set("second", 2)
	go bc.Set("third", 3)
	go bc.Remove("second")
	time.Sleep(time.Second)
	fmt.Println(bc.Get("first"))
}
