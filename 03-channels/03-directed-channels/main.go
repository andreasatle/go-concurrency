package main

import "fmt"

func generateMessage(ch chan<- int) {
	a, b := 0, 1
	for i := 0; i < 12; i++ {
		ch <- b
		a, b = b, a+b
	}
	close(ch)
}

func relayMessage(recv <-chan int, send chan<- int) {
	for r := range recv {
		send <- r
	}
	close(send)
}

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)
	go generateMessage(ch1)
	go relayMessage(ch1, ch2)

	for r := range ch2 {
		fmt.Println(r)
	}
}
