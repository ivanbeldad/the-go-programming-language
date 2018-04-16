// Write a version of rotate that operates in a single pass.

package main

import (
	"fmt"
	"math"
)

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Original:\t%v\n", s)
	Rotate(s, -3)
	fmt.Printf("Shifted:\t%v\n", s)
}

// Rotate move elements in the slice n positions to the right
// if it is positive, to the left otherwise
func Rotate(s []int, n int) {
	for n > len(s) {
		n -= len(s)
	}
	if n == 0 {
		return
	}
	absN := int(math.Abs(float64(n)))
	for i := 0; i < absN; i++ {
		if n > 0 {
			last := s[len(s)-1]
			for j := len(s) - 1; j > 0; j-- {
				s[j] = s[j-1]
			}
			s[0] = last
		} else {
			first := s[0]
			for j := 0; j < len(s)-1; j++ {
				s[j] = s[j+1]
			}
			s[len(s)-1] = first
		}
	}
}

// ReverseRotate move elements in the slice n positions
// to the right if it is positive, to the left otherwise
func ReverseRotate(s []int, n int) {
	for n > len(s) {
		n -= len(s)
	}
	if n == 0 {
		return
	}
	if n > 0 {
		reverse(s)
	}
	absN := int(math.Abs(float64(n)))
	reverse(s[:absN])
	reverse(s[absN:])
	if n < 0 {
		reverse(s)
	}
}

// ReverseNoLenRotate move elements in the slice n positions
// to the right if it is positive, to the left otherwise
func ReverseNoLenRotate(s []int, n int) {
	l := len(s)
	for n > l {
		n -= l
	}
	if n == 0 {
		return
	}
	if n > 0 {
		reverse(s)
	}
	absN := int(math.Abs(float64(n)))
	reverse(s[:absN])
	reverse(s[absN:])
	if n < 0 {
		reverse(s)
	}
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
