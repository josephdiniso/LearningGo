package main

import (
    "fmt"
    "sync"
    "sync/atomic"
)
/*
Atomic Counters

Used to increment or change a variable that other goroutines are also updating
at the same time. Without atomic counters the value would be corrupted
*/

func main() {

	// Counter var
    var ops uint64

	// Wait group to ensure all processes are finished
    var wg sync.WaitGroup

	// Creates fifty goroutines to increment counter 1000 times
    for i := 0; i < 50; i++ {
        wg.Add(1)
		
        go func() {
            for c := 0; c < 1000; c++ {
				// Uses atomic AddUint64 to increment by giving it
				// memory address and increment
                atomic.AddUint64(&ops, 1)
			}
			// Kills wg
            wg.Done()
        }()
    }
	// Waits until all processes are finished
    wg.Wait()

	// Prints counter value
	fmt.Println("ops:", ops)
	
	// Can use functions like atomic.LoadUint64 to view atomics while being updated

}