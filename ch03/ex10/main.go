// Write a non-recursive version of comma, using bytes.Buffer instead of string concatenation.

package main

import (
	"bytes"
	"fmt"
)

const (
	commaSpace = 3
)

func main() {
	fmt.Printf("Final result: %s\n", comma("12345678"))
	fmt.Printf("Final result: %s\n", comma("123456789"))
}

func comma(s string) string {
	finalLength := len(s)
	offset := finalLength % commaSpace
	if offset == 0 {
		offset = commaSpace
	}
	var buffer bytes.Buffer
	buffer.WriteString(s[:offset])
	for i := offset; i < finalLength; i += commaSpace {
		buffer.WriteString(",")
		buffer.WriteString(s[i : i+commaSpace])
	}
	return buffer.String()
}

func recursiveComma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return recursiveComma(s[:n-3]) + "," + s[n-3:]
}
