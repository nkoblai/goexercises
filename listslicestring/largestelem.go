package main

import (
	"fmt"
)

var list = []int{
	48, 96, 86, 68,
	57, 82, 63, 70,
	37, 112, 83, 27,
	19, 97, 112, -2,
}

// LargestElement finding largest element in list and returns it.
func LargestElement(list []int) int {
	maxVal := list[len(list)-1]
	for _, v := range list {
		if v >= maxVal {
			maxVal = v
		}
	}
	return maxVal
}

// Write a function that returns the largest element in a list.
func main() {
	fmt.Println(LargestElement(list))
}
