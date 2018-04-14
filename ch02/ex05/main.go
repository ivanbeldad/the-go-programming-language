// The expression x&(x-1) clears the rightmost non-zero bit of x. Write a version
// of PopCount that counts bits by using this fact, and assess its performance.

package main

import (
	"fmt"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	fmt.Printf("%d\n", PopCount(1255))
	fmt.Printf("%d\n", RightPopCount(1255))
}

// RightPopCount returns the population count (number of set bits) of x.
func RightPopCount(x uint64) int {
	total := 0
	for x != 0 {
		total++
		x = x & (x - 1)
	}
	return total
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// MyPopCount returns the population count (number of set bits) of x.
func MyPopCount(x uint64) int {
	var i uint8
	total := 0
	for i = 0; i < 64; i++ {
		total += int(x & 1)
		x = x >> 1
	}
	return total
}

// MyBreakPopCount returns the population count (number of set bits) of x.
func MyBreakPopCount(x uint64) int {
	var i uint8
	total := 0
	for i = 0; i < 64; i++ {
		total += int(x & 1)
		x = x >> 1
		if x == 0 {
			break
		}
	}
	return total
}

// MyBreakUnsignedPopCount returns the population count (number of set bits) of x.
func MyBreakUnsignedPopCount(x uint64) int {
	var i uint8
	var total uint64
	for i = 0; i < 64; i++ {
		total += x & 1
		x = x >> 1
		if x == 0 {
			break
		}
	}
	return int(total)
}

// LoopPopCount returns the population count (number of set bits) of x.
func LoopPopCount(x uint64) int {
	total := 0
	var i uint
	for i = 0; i <= 8; i++ {
		total += int(pc[byte(x>>(i*8))])
	}
	return total
}
