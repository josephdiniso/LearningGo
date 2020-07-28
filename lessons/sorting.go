package main

import (
    "fmt"
    "sort"
)

/*
To sort by function in Go, you need a corresponding type, byLength is just an 
alias for the builtin []string type
*/
type byLength []string

/*
Changes built in sort for type byLength to get the length of each item,
swap the two items, and then check if the length of the swapped is less than
the other
*/
func (s byLength) Len() int {
	return len(s)
}
func (s byLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	// Sorting is type specific, strings are sorted alphabetically, etc.

	// Sorting is in place and thus changes the original slices
    strs := []string{"c", "a", "b"}
    sort.Strings(strs)
    fmt.Println("Strings:", strs)

    ints := []int{7, 2, 4}
    sort.Ints(ints)
    fmt.Println("Ints:   ", ints)

	// Checks if a slice is already sorted
    s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
	
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(byLength(fruits))
	fmt.Println(fruits)
}