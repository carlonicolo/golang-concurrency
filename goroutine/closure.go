package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	// what is the output
	// TODO: fix the issue

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		// if you want that the goroutine works on a specific value
		// you need to pass it as param
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}
