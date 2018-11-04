package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	sayChao := func() {
		defer wg.Done()
		fmt.Println("chao!")
	}

	wg.Add(3)
	go sayHello(&wg)
	go func() {
		defer wg.Done()
		fmt.Println("hi!")
	}()
	go sayChao()

	// What a beautiful join!
	wg.Wait()

	// Using time.Sleep for printing something is just a conflict, not a join.
	//time.Sleep(1 * time.Second)
}

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("hello!")
}
