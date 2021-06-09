package main

import (
	"fmt"
	"sync"
)

// Use sync.WaitGroup to wait for go-routine to end. (like fork-join)
func main() {
	var data int
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		data++
	}()

	wg.Wait()
	fmt.Printf("The value of the data is %v\n", data)
	fmt.Println("Done")
}
