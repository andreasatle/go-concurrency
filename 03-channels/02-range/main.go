package main

import "fmt"

func main() {
	ch := make(chan int, 6)
	go func() {
		for i := 1; i <= 6; i++ {
			ch <- i
		}
		close(ch)
	}()
	for r := range ch {
		fmt.Println(r)
	}
}
