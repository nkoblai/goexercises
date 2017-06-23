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

// FindElement checks whether an element occurs in a list.
func FindElement(list []int, element int) bool {
	for _, v := range list {
		if v == element {
			return true
		}
	}
	return false
}

// Write a function that checks whether an element occurs in a list.
func main() {
	fmt.Println(FindElement(list, 2))
	fmt.Println(FindElement(list, 96))
}
