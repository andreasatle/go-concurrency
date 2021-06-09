package main

import (
	"fmt"
	"time"
)

func outputThreeTimes(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond)
	}
}

// Three different ways to start a go-routine.
// The third is kind of the same as the first...
func main() {
	outputThreeTimes("direct call")

	go outputThreeTimes("go-routine call")

	go func() {
		outputThreeTimes("Anonymous go-routine call")
	}()

	fv := outputThreeTimes
	go fv("go-routine with function value call")
	fmt.Println("Waiting for go-routines to execute.")
	time.Sleep(time.Second)
}
