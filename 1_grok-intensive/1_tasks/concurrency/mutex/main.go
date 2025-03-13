package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	c.value++
	c.mu.Unlock()
}

func main() {
	counter := Counter{}
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			counter.Increment()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Counter:", counter.value)
}
