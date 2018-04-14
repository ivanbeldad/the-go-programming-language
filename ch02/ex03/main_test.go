package main

import (
	"testing"
)

func BenchmarkPopcount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		PopCount(uint64(n))
	}
}

func BenchmarkLoopPopcount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		LoopPopCount(uint64(n))
	}
}
