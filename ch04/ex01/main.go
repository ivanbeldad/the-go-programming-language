// Write a function that counts the number of bits that are different in two SHA256
// hashes. (See PopCount from Section 2.6.2.)

package main

import (
	"fmt"
	"math/rand"
	"time"
)

var bitc = map[byte]int{
	1:   1,
	2:   1,
	3:   2,
	4:   1,
	5:   2,
	6:   2,
	7:   3,
	8:   1,
	9:   2,
	0xA: 2,
	0xB: 3,
	0xC: 2,
	0xD: 3,
	0xE: 3,
	0xF: 4,
}

func main() {
	rand.Seed(time.Now().UnixNano())
	s := [32]byte{}
	for i := range s {
		r := rand.Intn(256)
		s[i] = byte(r)
	}
	fmt.Printf("Result: %d\n", bitCount(s))
	fmt.Printf("Result: %d\n", crazyBitCount(s))
}

func bitCount(sha [32]byte) int {
	t := 0
	for i := range sha {
		for j := 0; j < 8; j++ {
			t += int(sha[i] & 1)
			sha[i] = sha[i] >> 1
		}
	}
	return t
}

func crazyBitCount(sha [32]byte) int {
	t := 0
	for i := range sha {
		t += bitc[sha[i]&15]
		t += bitc[sha[i]&240>>4]
	}
	return t
}
