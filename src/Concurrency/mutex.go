package main

import (
	"fmt"
	"sync"
	"time"
)

//Mutual exclusion is used to prevent two goroutines to access the
//same variable at the same time and avoid race conditions
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

//Create a funcion that executes mutex surroanding it bit lock and unlock methods
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
	//You can also defer your unlock to ensure the mutex willbe unlock once the function finishes
}

func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
