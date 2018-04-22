// Implement these additional methods:
// func (*IntSet) Len() int // return the number of elements
// func (*IntSet) Remove(x int) // remove x from the set
// func (*IntSet) Clear() // remove all elements from the set
// func (*IntSet) Copy() *IntSet // return a copy of the set

package main

import (
	"bytes"
	"fmt"
)

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// Len return the number of elements
func (s *IntSet) Len() (n int) {
	for i := 0; i < len(s.words); i++ {
		for bit := uint(0); bit < 64; bit++ {
			if s.words[i]&(1<<bit) != 0 {
				n++
			}
		}
	}
	return
}

// Remove remove x from the set
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] ^= 1 << bit
}

// Clear remove all elements from the set
func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] &= 0
	}
}

// Copy return a copy of the set
func (s IntSet) Copy() (rs IntSet) {
	rs.words = make([]uint64, len(s.words))
	copy(rs.words, s.words)
	return
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	x := IntSet{}
	x.Add(3)
	x.Add(30)
	x.Add(2)
	x.Add(5)

	fmt.Printf("\nAdding 3, 30, 2 and 5 to X\n")
	fmt.Printf("X: %s\tLen: %d\n", &x, x.Len())

	y := x.Copy()

	fmt.Printf("\nCopy X to Y\n")
	fmt.Printf("Y: %s\tLen: %d\n", &y, y.Len())

	x.Remove(3)

	fmt.Printf("\nRemoving number 3 in X\n")
	fmt.Printf("X: %s\tLen: %d\n", &x, x.Len())

	x.Clear()

	fmt.Printf("\nClear X\n")
	fmt.Printf("X: %s\tLen: %d\n", &x, x.Len())
	fmt.Printf("Y: %s\tLen: %d\n", &y, y.Len())
}
