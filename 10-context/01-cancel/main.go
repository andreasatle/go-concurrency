package main

import (
	"context"
	"fmt"
)

// Example using context WithCancel, printing 5 numbers
func main() {
	generator := func(ctx context.Context) <-chan int {

		ch := make(chan int)

		go func() {

			defer close(ch)
			for n := 1; ; n++ {
				select {
				case ch <- n:
				case <-ctx.Done():
					return
				}
			}
		}()

		return ch
	}

	ctx, cancel := context.WithCancel(context.Background())
	ch := generator(ctx)
	for num := range ch {
		fmt.Println(num)
		if num == 5 {
			cancel()
		}
	}
}
