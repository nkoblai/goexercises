package stringer

import (
	"testing"
)

func BenchmarkStrConvItoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IPAddr.Convert(IPAddr{127, 0, 0, 1})
	}
}

func BenchmarkFmtSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IPAddr.String(IPAddr{127, 0, 0, 1})
	}
}
