package main

import (
	"fmt"
	"strings"
)

// IsPalindrome tests whether a string is a palindrome.
func IsPalindrome(s string) bool {
	res := ""
	for i := range s {
		if s[i] >= 97 && s[i] <= 122 || s[i] >= 65 && s[i] <= 90 || s[i] >= 48 && s[i] <= 57 {
			res += string(s[i])
		}
	}
	res = strings.ToLower(res)
	for i := 0; i < len(res)/2; i++ {
		if res[i] != res[len(res)-i-1] {
			return false
		}
	}
	return true
}

// Write a function that tests whether a string is a palindrome.
func main() {
	fmt.Println(IsPalindrome("n1on"))
	fmt.Println(IsPalindrome("UFO tofu?"))
	fmt.Println(IsPalindrome("Yo, banana boy!"))
	fmt.Println(IsPalindrome("Borrow or rob?"))
}
