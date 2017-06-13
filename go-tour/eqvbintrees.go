package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Walker launches Walk in a new goroutine,
// and returns a read-only channel of values.
func Walker(t *tree.Tree) chan int {
	ch := make(chan int)
	go func() {
		Walk(t, ch)
		close(ch)
	}()
	return ch
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := Walker(t1), Walker(t2)

	for {

		v1, ok1 := <-ch1
		v2, ok2 := <-ch2

		if !ok1 || !ok2 {
			return ok1 == ok2
		}
		if v1 != v2 {
			break
		}
	}
	return false
}

// 1. Implement the Walk function.
// 2. Test the Walk function.
// The function tree.New(k) constructs a randomly-structured binary tree holding the values k, 2k, 3k, ..., 10k.
// Create a new channel ch and kick off the walker:
// go Walk(tree.New(1), ch)
// Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.
// 3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.
// 4. Test the Same function.
// Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.

func main() {
	ch := make(chan int)
	go Walk(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
		if i == 10 {
			close(ch)
		}
	}
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
