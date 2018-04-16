// Write an in-place function that squashes each run of adjacent Unicode spaces
// (see unicode.IsSpace) in a UTF-8-encoded []byte slice into a single ASCII space.

package main

import (
	"fmt"
	"unicode"
)

func main() {
	b := []byte("a  a    ad`    本a   1a")
	fmt.Printf("Original:\t%q\n", b)
	squashSpaces(&b)
	fmt.Printf("Modified:\t%q\n", b)
}

func squashSpaces(bytes *[]byte) {
	runes := []rune(string(*bytes))
	for i := 0; i < len(runes)-1; i++ {
		if unicode.IsSpace(runes[i]) && unicode.IsSpace(runes[i+1]) {
			runes = append(runes[:i], runes[i+1:]...)
			i--
		}
	}
	*bytes = []byte(string(runes))
}
