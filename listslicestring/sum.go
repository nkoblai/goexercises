package sum

import "fmt"

var list = []int{
	48, 96, 86, 68,
	57, 82, 63, 70,
	37, 112, 83, 27,
	19, 97, 112, -2,
}

// SumFor computes the sum of the numbers using for-loop.
func SumFor(list []int) int {
	res := 0
	for _, v := range list {
		res += v
	}
	return res
}

// SumWhile computes the sum using while-loop.
func SumWhile(list []int) int {
	res, i := 0, 0
	for i != len(list) {
		res += list[i]
		i++
	}
	return res
}

// SumRecursion computes the sum using recursion.
func SumRecursion(list []int) int {
	if len(list) == 1 {
		return list[0]
	}
	return list[0] + SumRecursion(list[1:])
}

// Write three functions that compute the sum of the numbers in a list.
func main() {
	fmt.Println(SumFor(list))
	fmt.Println(SumWhile(list))
	fmt.Println(SumRecursion(list))
}
