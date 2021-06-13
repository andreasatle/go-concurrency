package main

import (
	"context"
	"fmt"
	"time"
)

type data struct {
	result string
}

func compute(ctx context.Context) <-chan data {

	ch := make(chan data)

	go func() {
		defer close(ch)
		deadline, ok := ctx.Deadline()
		if ok {
			if deadline.Sub(time.Now().Add(50*time.Millisecond)) < 0 {
				fmt.Println("Process timed out")
				return
			}
		}
		time.Sleep(50 * time.Millisecond)
		select {
		case ch <- data{"123"}:
		case <-ctx.Done():
			return
		}
	}()

	return ch

}

func main() {
	runWithDeadline(100 * time.Millisecond)
	runWithDeadline(10 * time.Millisecond)
}

func runWithDeadline(duration time.Duration) {
	fmt.Println("Run with Deadline:", duration)
	deadline := time.Now().Add(duration)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	ch := compute(ctx)
	d, ok := <-ch
	if ok {
		fmt.Println("Work is complete:", d)
	}
}
