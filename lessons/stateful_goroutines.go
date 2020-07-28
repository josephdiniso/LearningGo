package main

import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

/*
Stateful Goroutines

Used as a built in alternative to mutexes
*/

// State is owned by a single goroutine and other goroutines will send messages
// to owner with corresponding replies
type readOp struct {
    key  int
    resp chan int
}
type writeOp struct {
    key  int
    val  int
    resp chan bool
}

func main() {
	// Counting number of performed operations
    var readOps uint64
    var writeOps uint64

	// Creates read write channels to be used by other goroutines for R/W
	// requests
    reads := make(chan readOp)
    writes := make(chan writeOp)

	// Creates parent function
    go func() {
        var state = make(map[int]int)
        for {
			// Responds to first R/W it receives and acts on it
            select {
            case read := <-reads:
                read.resp <- state[read.key]
            case write := <-writes:
                state[write.key] = write.val
                write.resp <- true
            }
        }
    }()
	// Creates 100 child instances
    for r := 0; r < 100; r++ {
        go func() {
            for {
				// Creates read operaton with data to send
                read := readOp{
                    key:  rand.Intn(5),
					resp: make(chan int)}
				// Sends read to read channel
				reads <- read
				// Waits for read response
                <-read.resp
                atomic.AddUint64(&readOps, 1)
                time.Sleep(time.Millisecond)
            }
        }()
    }
	// Creates 10 child instances to write
    for w := 0; w < 10; w++ {
        go func() {
            for {
                write := writeOp{
                    key:  rand.Intn(5),
                    val:  rand.Intn(100),
                    resp: make(chan bool)}
                writes <- write
                <-write.resp
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
}