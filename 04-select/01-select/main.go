package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch1 <- "one"
	}()

	go func() {
		time.Sleep(300 * time.Millisecond)
		ch1 <- "two"
	}()

	for i := 0; i < 2; i++ {
		select {
		case v := <-ch1:
			fmt.Println(v)
		case v := <-ch2:
			fmt.Println(v)
		}
	}

}
