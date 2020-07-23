package main

import (
    "fmt"
    "sync"
	"time"
)
/*
Wait Groups:

Used to block a program until all groups have returned 


*/
func worker(id int, wg *sync.WaitGroup, job int, results chan<- int) {
	// Notifies waitgroup when complete
	defer wg.Done()
	time.Sleep(time.Second)
	results <- job*10
	}


func main() {
	wg := new(sync.WaitGroup)
	numJobs := 5
	// Creates channel to receive values from worker
	results := make(chan int, numJobs)

	// Launches goroutines and increments the waitgroup counter each time
	s := []int{1,9,12,6,15}
	for i := 0; i < 5; i++ {
        wg.Add(1)
        go worker(i, wg, s[i], results)
	}
	// Waits until all workgroups are closed
	wg.Wait()

	// Iterates through results channel and prints values
	for i:=1; i<=numJobs; i++ {
		fmt.Println(<-results)
	}
	
	// Closes results channel
	close(results)


}