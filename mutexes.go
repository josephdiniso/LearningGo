package main

import (
    "fmt"
    "math/rand"
    "sync"
    "sync/atomic"
    "time"
)
/*
Mutexes

Used to safely access data across multiple goroutines
https://gobyexample.com/mutexes
*/
func main() {
	// State will be a map
    var state = make(map[int]int)

	// Mutex will synchronize access to state
    var mutex = &sync.Mutex{}

	// Atomic counters to count read and writes
    var readOps uint64
    var writeOps uint64

	// Creates 100 instances of a goroutine
    for r := 0; r < 100; r++ {
        go func() {
            total := 0
            for {
				// Creates key and locks mutex to read, state, and ensure
				// exclusive access
                key := rand.Intn(5)
                mutex.Lock()
                total += state[key]
				mutex.Unlock()
				// Iterates for one read
                atomic.AddUint64(&readOps, 1)

                time.Sleep(time.Millisecond)
            }
        }()
    }
	// Runs 10 goroutines to write to state in the same manner
    for w := 0; w < 10; w++ {
        go func() {
            for {
                key := rand.Intn(5)
                val := rand.Intn(100)
                mutex.Lock()
                state[key] = val
                mutex.Unlock()
                atomic.AddUint64(&writeOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }

    time.Sleep(time.Second)

    readOpsFinal := atomic.LoadUint64(&readOps)
    fmt.Println("readOps:", readOpsFinal)
    writeOpsFinal := atomic.LoadUint64(&writeOps)
    fmt.Println("writeOps:", writeOpsFinal)

	// Locks mutex to read final state
    mutex.Lock()
    fmt.Println("state:", state)
    mutex.Unlock()
}