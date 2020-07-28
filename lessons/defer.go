package main

import (
    "fmt"
    "os"
)

/*
Defer is used to ensure a function call is performed later in a program's execution
Usually for cleanup

Runs the function at retunr
*/
func main() {

	f := createFile("/tmp/defer.txt")
	// Allows us to close file after writing to file
    defer closeFile(f)
    writeFile(f)
}

func createFile(p string) *os.File {
    fmt.Println("creating")
    f, err := os.Create(p)
    if err != nil {
        panic(err)
    }
    return f
}

func writeFile(f *os.File) {
    fmt.Println("writing")
    fmt.Fprintln(f, "data")

}

func closeFile(f *os.File) {
	fmt.Println("closing")
	// Important to check for errors even in a deferred function
    err := f.Close()

    if err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}