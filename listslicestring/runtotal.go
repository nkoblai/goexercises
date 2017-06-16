package main

import "fmt"

var list = []int{
	48, 96, 86, 68,
	57, 82, 63, 70,
	37, 112, 83, 27,
	19, 97, 112, -2,
}

// RuningTotal computes the running total of a list.
func RuningTotal(list []int) int {
	res := 0
	for _, v := range list {
		res += v
	}
	return res
}

// Write a function that computes the running total of a list.
func main() {
	fmt.Println(RuningTotal(list))
}
