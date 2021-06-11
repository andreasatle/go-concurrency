package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func generator(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, num := range nums {
			select {
			case out <- num:
			case <-done:
				return
			}
		}
	}()
	return out
}

func square(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for num := range in {
			select {
			case out <- num * num:
			case <-done:
				return
			}
		}
	}()

	return out
}

func main() {
	nums := createSequence(60)
	squareWithFan(nums...)

	// Insert delay before check of number of go-routines active.
	time.Sleep(10 * time.Millisecond)
	fmt.Printf("Number of go-routines active after call to squareWithFan: %d\n", runtime.NumGoroutine())
}

func createSequence(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}
	return nums
}

func squareWithFan(nums ...int) {
	done := make(chan struct{})
	defer close(done)

	nProcs := runtime.GOMAXPROCS(runtime.NumCPU())
	in := generator(done, nums...)

	cnt := 0
	for num := range fanIn(done, fanOut(done, in, nProcs)) {
		fmt.Printf("%d ", num)
		cnt++
		// Stop after 10 values (not necessarily the first ten...)
		if cnt >= 10 {
			break
		}
	}
	fmt.Println()
	fmt.Printf("Number of go-routines active before close(done): %d\n", runtime.NumGoroutine())

}

// Create n channels to split the work
func fanOut(done <-chan struct{}, in <-chan int, n int) []<-chan int {
	outs := make([]<-chan int, n)
	for i := range outs {
		outs[i] = square(done, in)
	}
	return outs
}

// Collect the n channels into one.
// Observe that we need an extra go-routine with the close-statement.
func fanIn(done <-chan struct{}, chs []<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	collect := func(ch <-chan int) {
		defer wg.Done()
		for num := range ch {
			select {
			case out <- num:
			case <-done:
				return
			}

		}
	}

	wg.Add(len(chs))
	for _, ch := range chs {
		go collect(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
