package main

import "fmt"

var list = []int{
	1, 2, 3, 4, 5,
}

// Reverse takes slice of integers and returns same reversed slice.
func Reverse(list []int) []int {
	for i := len(list)/2 - 1; i >= 0; i-- {
		k := len(list) - 1 - i
		list[i], list[k] = list[k], list[i]
	}
	return list
}

// Write function that reverses a list, preferably in place.
func main() {
	fmt.Println(Reverse(list))
}
