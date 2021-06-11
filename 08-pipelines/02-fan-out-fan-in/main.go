package main

import (
	"fmt"
	"runtime"
	"sync"
)

func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for num := range in {
			out <- num * num
		}
		close(out)
	}()
	return out

}

func main() {
	nums := createSequence(40)
	squareWithFan(nums...)
}

func createSequence(n int) []int {
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		nums[i] = i
	}
	return nums
}

func squareWithFan(nums ...int) {
	nProcs := runtime.GOMAXPROCS(runtime.NumCPU())
	in := generator(nums...)

	for num := range fanIn(fanOut(in, nProcs)) {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
}

// Create n channels to split the work
func fanOut(in <-chan int, n int) []<-chan int {
	outs := make([]<-chan int, n)
	for i := range outs {
		outs[i] = square(in)
	}
	return outs
}

// Collect the n channels into one.
// Observe that we need an extra go-routine with the close-statement.
func fanIn(chs []<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	collect := func(ch <-chan int) {
		for num := range ch {
			out <- num
		}
		wg.Done()
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
