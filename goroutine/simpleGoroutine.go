package main

import (
	"fmt"
	"sync"
)

func main() {
	// TODO: modify the program
	// to print the value data as 1
	// deterministically

	var data int
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		data++
	}()

	wg.Wait()

	fmt.Printf("The value of data is %v\n", data)
	fmt.Println("Done...")
}
