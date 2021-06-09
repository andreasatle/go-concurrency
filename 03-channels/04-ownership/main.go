package main

import "fmt"

func main() {
	// owner is responsible for creating the channel, feeding it and closing it when done.
	// This design pattern is very common.
	owner := func() <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := 0; i < 5; i++ {
				ch <- i + 1
			}
		}()
		return ch
	}
	consumer := func(ch <-chan int) {
		for r := range ch {
			fmt.Println("Received:", r)
		}
		fmt.Println("Done receiving!")
	}

	ch := owner()
	consumer(ch)
}
