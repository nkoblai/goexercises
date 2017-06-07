package main

import "golang.org/x/tour/pic"

// Pic returns a slice of length dy, each element of which is a slice of dx 8-bit
// unsigned integers
func Pic(dx, dy int) [][]uint8 {
	x := make([][]uint8, dy)
	for i := range x {
		x[i] = make([]uint8, dx)
		for z := range x[i] {
			x[i][z] = uint8(i ^ z)
		}
	}
	return x
}

// Implement Pic. It should return a slice of length dy, each element of which is
// a slice of dx 8-bit unsigned integers. When you run the program, it will display
// your picture, interpreting the integers as grayscale (well, bluescale) values.
func main() {
	pic.Show(Pic)
}
