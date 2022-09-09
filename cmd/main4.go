package main

import (
	"fmt"
	"sync"
	"time"
)

type Cashe struct {
	storage map[string]int
	mu      sync.RWMutex
}

func (c *Cashe) Increase(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] += value
}

func (c *Cashe) Set(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.storage[key] = value
}

func (c *Cashe) Get(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.storage[key]
}

func (c *Cashe) Remove(key string) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	delete(c.storage, key)

}

const (
	k1   = "key1"
	step = 7
)

func main() {
	cashe := Cashe{storage: make(map[string]int)}
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			cashe.Increase(k1, step)
			time.Sleep(time.Millisecond)
		}()
		wg.Wait()
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			cashe.Set(k1, step*i)
			time.Sleep(time.Millisecond)
		}()
		wg.Wait()
	}

	fmt.Println(cashe.Get(k1))

}
