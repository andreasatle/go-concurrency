package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	withRaceCondition()
	withoutRaceCondition()
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}

func withRaceCondition() {
	start := time.Now()
	var t *time.Timer
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		t.Reset(randomDuration())
	})

	time.Sleep(5 * time.Second)
}

func withoutRaceCondition() {
	start := time.Now()
	var t *time.Timer
	ch := make(chan struct{})
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Now().Sub(start))
		ch <- struct{}{}
	})

	for time.Since(start) < 5*time.Second {
		<-ch
		t.Reset(randomDuration())
	}
}
