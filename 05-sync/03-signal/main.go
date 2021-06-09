package main

import (
	"fmt"
	"sync"
)

var sharedRsc = make(map[string]interface{})

func main() {
	var wg sync.WaitGroup

	// Init conditional variable c
	mut := sync.Mutex{}
	c := sync.NewCond(&mut)

	// go-routine for data read
	wg.Add(1)
	go func() {
		defer wg.Done()

		// Step 1 - Apply lock
		c.L.Lock()

		// Step 2 - Wait for signal from other go-routine
		for _, ok := sharedRsc["rsc1"]; !ok; _, ok = sharedRsc["rsc1"] {
			c.Wait()
		}

		// Step 3 - Process data
		fmt.Println(sharedRsc["rsc1"])

		// Step 4 - Unlock
		c.L.Unlock()
	}()

	// go-routine for data write
	wg.Add(1)
	go func() {
		defer wg.Done()

		// Step 1 - Lock
		c.L.Lock()

		// Step 2 - Create data
		sharedRsc["rsc1"] = "foo"

		// Step 3 - Signal that data can be used
		c.Signal()

		// Step 4 Unlock
		c.L.Unlock()
	}()

	wg.Wait()
}
