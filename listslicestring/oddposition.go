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

// OddPosition returns slice with the elements on odd positions in given slice.
func OddPosition(list []int) []int {
	res := make([]int, 0)
	for i, v := range list[1:] {
		if i%2 == 0 {
			res = append(res, v)
		}
	}
	return res
}

// Write a function that returns the elements on odd positions in a list.
func main() {
	fmt.Println(OddPosition(list))
}
