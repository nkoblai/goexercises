package main

import (
	"fmt"
)

var list = []int{
	1, 2, 3, 4, 5,
}

// Reverse takes slice of integers and returns same reversed slice.
func Reverse(list []int) []int {
	if len(list) == 0 {
		return list
	}
	return append(Reverse(list[1:]), list[0])
}

// Write function that reverses a list, preferably in place.
func main() {
	fmt.Println(Reverse(list))
}
