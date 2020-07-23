package main

import (
    "fmt"
    "time"
)

func adder(id int, jobs <-chan int, results chan<- int) {
	// What happens when iterating over buffered channel with no items (?)
	for j := range jobs {
		fmt.Println("worker", id, "started job", <-jobs)
		time.Sleep(time.Second)
		// fmt.Println("worker", id, "finished job", j)
		results <- j + 7
	}
}


func main() {
	const numJobs = 5

	// Creates channel to send values to worker
	jobs := make(chan int, numJobs)
	// Creates channel to receive values from worker
	results := make(chan int, numJobs)

	// Creates worker pool of size 3
	for n := 1; n<=3; n++ {
		go adder(n, jobs, results)
	}

	// Sends 5 jobs to workers
	s := []int{1,5,7, 14, 18}
	for j := range(s) {
		jobs <- j
	}

	// Receives results from buffered channel
	for a := 1; a<= numJobs; a++ {
		fmt.Println(<-results)
	}
}