package main

import "fmt"

func main() {
	// Create a channel
	ch := make(chan int)

	go func(a, b int) {
		// Send a+b to channel
		ch <- a + b
	}(1, 2)

	// Receive r from channel
	r := <-ch

	fmt.Println("Computed value:", r)
}
