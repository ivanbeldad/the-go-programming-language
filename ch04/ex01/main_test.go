package main

import (
	"math/rand"
	"testing"
	"time"
)

var s = [32]byte{}

func init() {
	rand.Seed(time.Now().UnixNano())
	for i := range s {
		r := rand.Intn(256)
		s[i] = byte(r)
	}
}

func BenchmarkBitCount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		bitCount(s)
	}
}

func BenchmarkCrazyBitCount(b *testing.B) {
	for n := 0; n < b.N; n++ {
		crazyBitCount(s)
	}
}
