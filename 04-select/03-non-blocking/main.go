package main

import (
	"fmt"
	"time"
)

func main() {
	n := 5
	ch := make(chan string, n)

	go func() {
		for i := 0; i < n; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- "msg from go-routine"
		}
	}()

	// Wait for n messages from the channel.
	// Meanwhile process some batch job.
	for i, j := 0, 0; i < n; {
		select {
		case msg := <-ch:
			fmt.Println(msg)
			i++
		default:
			fmt.Printf("Working on batch %d while waiting...\n", j)
			time.Sleep(77 * time.Millisecond)
			j++
		}
	}
}
