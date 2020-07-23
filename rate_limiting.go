package main

import (
    "fmt"
    "time"
)
/*
Rate Limiting:

Mechanism for controlling resource utilization and maintaining quality of 
service.
https://gobyexample.com/rate-limiting
*/
func main() {

	// Channel with buffer size: 5
	requests := make(chan int, 5)
	// Sends 1-5 to requests channel
    for i := 1; i <= 5; i++ {
        requests <- i
	}
	// Closes channel
    close(requests)

	// Limiter channel receives a value every 200 ms
    limiter := time.Tick(200 * time.Millisecond)

	// Iterates through requests channel and waits 200ms before next iteration
    for req := range requests {
        <-limiter
        fmt.Println("request", req, time.Now())
    }

	// A bursy limiter allows short bursts of requests while preserving overall rate limit
	// Allows for a burst of up to 3 events
    burstyLimiter := make(chan time.Time, 3)

	// Iterates three times and sends time 
    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }

	// Every 200 ms. tries to add a value to the buffered limitered channel
	// Up to three events can be added
    go func() {
        for t := range time.Tick(200 * time.Millisecond) {
            burstyLimiter <- t
        }
    }()
	
	// Fills bursty requests channel
    burstyRequests := make(chan int, 5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i
    }
	close(burstyRequests)
	// Serves first three requests immediately because burstyLimiter is buffered 
	// and nonblocking, then serves every 200 ms.
    for req := range burstyRequests {
		/*
		Because buffered channels are nonblocking when they have entries and blocking
		when they have no entries, the loop will only iterate when burstyLimiter
		receives a value
		*/
		<-burstyLimiter
        fmt.Println("request", req, time.Now())
    }
}