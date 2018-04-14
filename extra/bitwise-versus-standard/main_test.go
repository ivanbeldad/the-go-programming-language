package bitwiseversusstandard

import (
	"testing"
)

func BenchmarkStandard(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isEvenStandard(i)
	}
}

func BenchmarkBitwise(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isEvenBitwise(i)
	}
}
