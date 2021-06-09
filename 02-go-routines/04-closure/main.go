package main

import (
	"fmt"
	"sync"
)

// Without arguments to go-routine, the program outputs 4, 4, 4
// The result should be (a permutation of) 1, 2, 3

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}
