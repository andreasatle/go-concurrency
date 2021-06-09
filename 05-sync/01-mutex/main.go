package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Add mutex to get thread safe updates of the balance
func main() {
	n := 100
	runtime.GOMAXPROCS(4)
	balance := 0
	var wg sync.WaitGroup
	var mut sync.Mutex // Add a mutex lock

	deposit := func(amount int) {
		mut.Lock()         // lock before deposit
		defer mut.Unlock() // Unlock after deposit
		balance += amount
	}

	withdrawal := func(amount int) {
		mut.Lock()         // Lock before withdrawal
		defer mut.Unlock() // Unlock after withdrawal
		balance -= amount
	}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			deposit(1)
		}()
	}

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			withdrawal(1)
		}()
	}
	wg.Wait()
	fmt.Println(balance)
}
