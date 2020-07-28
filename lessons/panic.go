package main

import "os"

/*
A panic means something went wrong and is used to handle unexpected errors
*/
func main() {

    panic("a problem")

	// Program causes panic, print error message and goroutine traces, and exit
	// with nonzero status
    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }
}