// Write a function that reports whether two strings are anagrams of each other,
// that is, they contain the same letters in a different order.

package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Printf("%s = %s\t\tResult: %t\n", "asdf", "dsaf", anagram("asdf", "dsaf"))
	fmt.Printf("%s = %s\t\tResult: %t\n", "asdfa", "dsaf", anagram("asdfa", "dsaf"))
	fmt.Printf("%s = %s\tResult: %t\n", "1112121", "2112111", anagram("1112121", "2112111"))
	fmt.Printf("%s = %s\tResult: %t\n", "界a世a界", "世aa界界", anagram("界a世a界", "世aa界界"))
	fmt.Printf("%s = %s\tResult: %t\n", "界a世a界", "世aa世界", anagram("界a世a界", "世aa世界"))
}

func anagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	r1 := []rune(s1)
	r2 := []rune(s2)
	for _, r := range r1 {
		byteIndex := bytes.IndexRune([]byte(string(r2)), r)
		if byteIndex == -1 {
			return false
		}
		i := utf8.RuneCountInString(s2[:byteIndex])
		r2[i] = 0
	}
	return true
}
