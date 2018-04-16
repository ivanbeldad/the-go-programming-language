package main

import "testing"

var s = []int{1, 2, 3, 4, 5}

var sLong []int

func init() {
	for i := 1; i <= 100; i++ {
		sLong = append(sLong, i)
	}
}

func BenchmarkReverse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Rotate(s, n)
	}
}

func BenchmarkReverseNoLen(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Rotate(sLong, n)
	}
}

func BenchmarkReverseRotate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReverseRotate(s, n)
	}
}

func BenchmarkReverseNoLenRotate(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReverseNoLenRotate(s, n)
	}
}

func BenchmarkReverseRotateLong(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReverseRotate(sLong, n)
	}
}

func BenchmarkReverseNoLenRotateLong(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ReverseNoLenRotate(sLong, n)
	}
}
