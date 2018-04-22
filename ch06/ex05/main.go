// The type of each word used by IntSet is uint64, but 64-bit arithmetic maybe
// inefficient on a 32-bit platform. Modify the program to use the uint type, which is the most
// efficient unsigned integer type for the platform. Instead of dividing by 64, define a constant
// holding the effective size of uint in bits, 32 or 64. You can use the perhaps too-clever
// expression 32 << (^uint(0) >> 63) for this purpose.

package main

import (
	"bytes"
	"fmt"
)

const is64 = uint64(^uint(0)) == ^uint64(0)

// IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet interface {
	Elems() []int
	Has(int) bool
	Add(int)
	AddAll(...int)
	Len() int
	Remove(int)
	Copy() IntSet
	Clear()
	String() string
}

// An IntSet32 is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet32 struct {
	words []uint32
}

// An IntSet64 is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet64 struct {
	words []uint64
}

// NewIntSet create a IntSet of 32 or 64 depending on the machine.
func NewIntSet() (i IntSet) {
	if is64 {
		fmt.Println("Creating IntSet64")
		i = &IntSet64{}
	} else {
		fmt.Println("Creating IntSet32")
		i = &IntSet32{}
	}
	return
}

// Elems returns a slice containing the elements of the set.
func (s *IntSet64) Elems() (e []int) {
	max := (len(s.words) * 64) - 1
	for i := 0; i <= max; i++ {
		if s.Has(i) {
			e = append(e, i)
		}
	}
	return
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet64) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet64) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll adds the non-negative values x to the set.
func (s *IntSet64) AddAll(x ...int) {
	for _, e := range x {
		s.Add(e)
	}
}

// Len return the number of elements
func (s *IntSet64) Len() (n int) {
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
func (s *IntSet64) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] ^= 1 << bit
}

// Clear remove all elements from the set
func (s *IntSet64) Clear() {
	for i := range s.words {
		s.words[i] &= 0
	}
}

// Copy return a copy of the set
func (s *IntSet64) Copy() (rs IntSet) {
	newSet := IntSet64{
		words: make([]uint64, len(s.words)),
	}
	copy(newSet.words, s.words)
	return &newSet
}

// UnionWith sets s to the union of s and t.
func (s *IntSet64) String() string {
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

// Elems returns a slice containing the elements of the set.
func (s *IntSet32) Elems() (e []int) {
	max := (len(s.words) * 32) - 1
	for i := 0; i <= max; i++ {
		if s.Has(i) {
			e = append(e, i)
		}
	}
	return
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet32) Has(x int) bool {
	word, bit := x/32, uint(x%32)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet32) Add(x int) {
	word, bit := x/32, uint(x%32)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// AddAll adds the non-negative values x to the set.
func (s *IntSet32) AddAll(x ...int) {
	for _, e := range x {
		s.Add(e)
	}
}

// Len return the number of elements
func (s *IntSet32) Len() (n int) {
	for i := 0; i < len(s.words); i++ {
		for bit := uint(0); bit < 32; bit++ {
			if s.words[i]&(1<<bit) != 0 {
				n++
			}
		}
	}
	return
}

// Remove remove x from the set
func (s *IntSet32) Remove(x int) {
	word, bit := x/32, uint(x%32)
	s.words[word] ^= 1 << bit
}

// Clear remove all elements from the set
func (s *IntSet32) Clear() {
	for i := range s.words {
		s.words[i] &= 0
	}
}

// Copy return a copy of the set
func (s *IntSet32) Copy() (rs IntSet) {
	newSet := IntSet32{
		words: make([]uint32, len(s.words)),
	}
	copy(newSet.words, s.words)
	return &newSet
}

// UnionWith sets s to the union of s and t.
func (s *IntSet32) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 32; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 32*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

func main() {
	x := NewIntSet()
	x.AddAll()
	x.AddAll(3, 30, 2, 5)

	fmt.Printf("\nAdding 3, 30, 2 and 5 to X\n")
	fmt.Printf("X: %s\tLen: %d\n", x, x.Len())

	y := x.Copy()

	fmt.Printf("\nCopy X to Y\n")
	fmt.Printf("Y: %s\tLen: %d\n", y, y.Len())

	x.Remove(3)

	fmt.Printf("\nRemoving number 3 in X\n")
	fmt.Printf("X: %s\tLen: %d\n", x, x.Len())

	x.Clear()

	fmt.Printf("\nClear X\n")
	fmt.Printf("X: %s\t\tLen: %d\n", x, x.Len())
	fmt.Printf("Y: %s\tLen: %d\n", y, y.Len())

	fmt.Printf("\nElems of Y\n")
	fmt.Printf("Y: %v\n", y.Elems())
}
