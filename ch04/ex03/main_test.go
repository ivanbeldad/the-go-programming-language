package main

import "testing"

var s = []int{1, 2, 3, 4, 5}
var a = [5]int{1, 2, 3, 4, 5}

var sLong = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
var aLong = [20]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

func BenchmarkSlice(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sliceReverse(s)
	}
}

func BenchmarkArray(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arrayReverse(&a)
	}
}

func BenchmarkSliceLong(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sliceReverse(sLong)
	}
}

func BenchmarkArrayLong(b *testing.B) {
	for n := 0; n < b.N; n++ {
		arrayReverseLong(&aLong)
	}
}
