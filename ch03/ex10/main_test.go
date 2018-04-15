package main

import "testing"

func BenchmarkRecursiveShort(b *testing.B) {
	s := "54326"
	for n := 0; n < b.N; n++ {
		recursiveComma(string(s))
	}
}

func BenchmarkIterativeShort(b *testing.B) {
	s := "54326"
	for n := 0; n < b.N; n++ {
		comma(string(s))
	}
}

func BenchmarkRecursiveLong(b *testing.B) {
	s := "34905645396745905645289634593459807634590876435890764386478675856797859789678967899756086"
	for n := 0; n < b.N; n++ {
		recursiveComma(string(s))
	}
}

func BenchmarkIterativeLong(b *testing.B) {
	s := "34905645396745905645289634593459807634590876435890764386478675856797859789678967899756086"
	for n := 0; n < b.N; n++ {
		comma(string(s))
	}
}
