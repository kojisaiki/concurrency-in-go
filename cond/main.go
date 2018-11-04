package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration, i int) {
		time.Sleep(delay)
		c.L.Lock()
		fmt.Println("-- lock for remove: ", i)
		queue = queue[1:]
		fmt.Println("Removed from queue: ", i)
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			fmt.Println("- main wait: ", i)
			c.Wait()
			fmt.Println("- main restart: ", i)
		}
		fmt.Println("Adding to queue: ", i)
		queue = append(queue, struct{}{})
		go removeFromQueue(1*time.Second, i)
		c.L.Unlock()
	}
}
