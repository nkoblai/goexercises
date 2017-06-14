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
	maxVal := 0
	for i := 0; i < len(list); i++ {
		for _, v := range list {
			if list[i] >= v {
				maxVal = list[i]
			}
		}
	}
	return maxVal
}

// Write a function that returns the largest element in a list.
func main() {
	fmt.Println(LargestElement(list))
}
