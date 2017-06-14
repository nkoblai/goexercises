package main

import (
	"fmt"
)

// IsPalindrome tests whether a string is a palindrome.
func IsPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] == s[len(s)-i-1] {
			return true
		}
	}
	return false
}

// Write a function that tests whether a string is a palindrome.
func main() {
	fmt.Println(IsPalindrome("test false"))
	fmt.Println(IsPalindrome("lol"))
	fmt.Println(IsPalindrome("noon"))
}
