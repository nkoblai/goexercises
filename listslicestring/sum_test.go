package sum

import (
	"testing"
)

func BenchmarkSumFor(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumFor(list)
	}
}

func BenchmarkSumWhile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumWhile(list)
	}
}

func BenchmarkSumRecursion(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SumRecursion(list)
	}
}
