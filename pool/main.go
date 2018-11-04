package main

import (
	"fmt"
	"sync"
)

func main() {
	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() interface{} {
			numCalcsCreated += 1
			mem := make([]byte, 1024)
			return &mem
		},
	}

	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := numWorkers; i > 0; i-- {
		go func(i int) {
			defer wg.Done()

			// If there no array in the pool(when return to pool is late), create new calc.
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)

			// Delay returning byte array to pool will occur creating new calc.
			// time.Sleep(1 * time.Microsecond)
		}(i)
	}

	wg.Wait()
	fmt.Printf("%d calculators were created.", numCalcsCreated)
}
