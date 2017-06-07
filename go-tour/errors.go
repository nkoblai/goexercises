package main

import (
	"fmt"
)

// ErrNegativeSqrt represents negative number.
type ErrNegativeSqrt float64

// Sqrt returns a square root of number using Newton's method.
func Sqrt(x float64) (float64, error) {
	z := 1.0
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	for n := 0; n <= 10; n++ {
		z = z - ((z*z - x) / (2 * z))
	}
	return z, nil
}

// ErrNegativeSqrt returns non-nil error value when given a negative number.
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("Cannot Sqrt negative number: ", float64(e))
}

// Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.
func main() {
	if _, err := Sqrt(-1); err != nil {
		fmt.Println(err)
	}
	v, _ := Sqrt(1)
	fmt.Println(v)
}
