package main

import "testing"

var s = []byte("Hello 本 World!")

var sLong = []byte("Hello 本 World! Hello 本 World! Hello 本 World!Hello 本 World! " +
	"Hello 本 World! Hello 本 World! Hello 本 World! Hello 本 World! Hello 本 World!" +
	"Hello 本 World! Hello 本 World! Hello 本 World! Hello 本 World! Hello 本 World!")

func BenchmarkReverse(b *testing.B) {
	for n := 0; n < b.N; n++ {
		reverse(&s)
	}
}

func BenchmarkReverseBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		reverseBuffer(&s)
	}
}

func BenchmarkReverseBuilder(b *testing.B) {
	for n := 0; n < b.N; n++ {
		reverseBuilder(&s)
	}
}

func BenchmarkReverseCopy(b *testing.B) {
	for n := 0; n < b.N; n++ {
		reverseCopy(&s)
	}
}

func BenchmarkReverseLong(b *testing.B) {
	for n := 0; n < b.N; n++ {
		reverse(&sLong)
	}
}

func BenchmarkReverseBufferLong(b *testing.B) {
	for n := 0; n < b.N; n++ {
		reverseBuffer(&sLong)
	}
}
func BenchmarkReverseBuilderLong(b *testing.B) {
	for n := 0; n < b.N; n++ {
		reverseBuilder(&sLong)
	}
}

func BenchmarkReverseCopyLong(b *testing.B) {
	for n := 0; n < b.N; n++ {
		reverseCopy(&sLong)
	}
}
