package main

import (
	"fmt"
	"math"
)

//Sqrt returns a square root of number using Newton's method.
func Sqrt(x float64) float64 {
	z := 1.0
	for n := 0; n <= 10; n++ {
		z = z - ((z*z - x) / (2 * z))
	}
	return z
}

//Implement the square root function using Newton's method
func main() {
	fmt.Println(Sqrt(2), Sqrt(4), Sqrt(10))
	fmt.Println(math.Sqrt(2), Sqrt(4), Sqrt(10))
}
