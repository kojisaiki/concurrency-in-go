package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("refer for loop scope")

	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greeting", "good bye"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
	// result by refer for loop scope:
	// good bye
	// good bye
	// good bye

	fmt.Println("")
	fmt.Println("refer functiion called scope")

	var wg2 sync.WaitGroup
	for _, salutation := range []string{"hello", "greeting", "good bye"} {
		wg2.Add(1)
		go func(salutation string) {
			defer wg2.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg2.Wait()
}
