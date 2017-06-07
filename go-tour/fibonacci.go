package main

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		defer func() {
			a, b = b, a+b
		}()
		return a
	}
}

// Implement a fibonacci function that returns a function (a closure)
// that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
