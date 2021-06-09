package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	runtime.GOMAXPROCS(4)

	counter := uint64(0)
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				//counter++

				// Replace counter++ with an atomic update
				atomic.AddUint64(&counter, 1)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("Counter:", counter)
}
