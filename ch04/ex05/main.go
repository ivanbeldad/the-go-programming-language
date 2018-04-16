// Write an in-place function to eliminate adjacent duplicates in a []string slice.

package main

import (
	"fmt"
)

func main() {
	strings := []string{"hello", "hello", "hello", "world", "world"}
	fmt.Printf("Original strings:\t%s\n", strings)
	removeDuplicates(&strings)
	fmt.Printf("Modified strings:\t%s\n", strings)
}

func removeDuplicates(strings *[]string) {
	ss := *strings
	removed := 0
	for i := 0; i < len(ss)-1; i++ {
		if ss[i] == ss[i+1] {
			ss = append(ss[0:i], ss[i+1:len(ss)]...)
			removed++
			i--
		}
	}
	*strings = ss // this works
	// strings = &ss // this doesn't work
}
