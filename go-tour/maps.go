package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

// WordCount returns a map of the counts of each “word” in the string s
func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	count := make(map[string]int)
	for word := range words {
		count[words[word]]++
	}
	return count
}

// Implement WordCount. It should return a map of the counts of each “word” in
// the string s. The wc.Test function runs a test suite against the provided
// function and prints success or failure.
func main() {
	wc.Test(WordCount)
}
