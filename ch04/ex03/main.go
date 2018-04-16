// Rewrite reverse to use an array pointer instead of a slice.

package main

import (
	"fmt"
)

const (
	arrL = 5
)

func main() {
	nums := []int{
		1, 2, 3, 4, 5,
	}
	nums2 := [arrL]int{
		1, 2, 3, 4, 5,
	}
	sliceReverse(nums)
	arrayReverse(&nums2)
	fmt.Printf("Slice: %v\n", nums)
	fmt.Printf("Array: %v\n", nums2)
}

func arrayReverse(s *[arrL]int) {
	for i, j := 0, arrL-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func arrayReverseLong(s *[20]int) {
	for i, j := 0, arrL-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func sliceReverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
