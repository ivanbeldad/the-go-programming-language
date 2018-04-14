package main

import (
	"testing"
)

func BenchmarkRightPopcount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		RightPopCount(uint64(n))
	}
}
func BenchmarkPopcount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PopCount(uint64(n))
	}
}

func BenchmarkMyPopcount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MyPopCount(uint64(n))
	}
}

func BenchmarkMyBreakPopCount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MyBreakPopCount(uint64(n))
	}
}

func BenchmarkMyBreakUnsignedPopCount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		MyBreakUnsignedPopCount(uint64(n))
	}
}

func BenchmarkLoopPopcount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LoopPopCount(uint64(n))
	}
}
