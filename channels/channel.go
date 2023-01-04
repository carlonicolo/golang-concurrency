package main

import "fmt"

func main() {
	ch := make(chan int)
	go func(a, b int) {
		c := a + b
		ch <- c
	}(1, 2)
	// TODO: get the value computed from gorouotine
	r := <-ch
	// fmt.Printf("Computed value %v\n", c)
	fmt.Printf("Computed value %v\n", r)

}
