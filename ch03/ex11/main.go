// Enhance comma so that it deals correctly with floating-point numbers and an optional sign.

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Original: %s\t\t Result: %s\n", "34567", comma("34567"))
	fmt.Printf("Original: %s\t Result: %s\n", "-34567", comma("-34567"))
	fmt.Printf("Original: %s\t Result: %s\n", "563767876", comma("563767876"))
	fmt.Printf("Original: %s\t Result: %s\n", "-563767876", comma("-563767876"))
	fmt.Printf("Original: %s\t\t Result: %s\n", "54", comma("54"))
	fmt.Printf("Original: %s\t\t Result: %s\n", "-54", comma("-54"))
	fmt.Printf("Original: %s\t Result: %s\n", "345.67", comma("345.67"))
	fmt.Printf("Original: %s\t Result: %s\n", "56767.6", comma("56767.6"))
	fmt.Printf("Original: %s\t Result: %s\n", "563767.876", comma("563767.876"))
	fmt.Printf("Original: %s\t Result: %s\n", "1563767.876", comma("1563767.876"))

}

func comma(s string) string {
	length := len(s)
	decimal := strings.IndexRune(s, '.')
	if decimal == -1 {
		decimal = 0
	} else {
		decimal = length - decimal
	}
	offset := 0
	if s[0] == '-' {
		offset++
	}
	if length <= 3+offset+decimal {
		return s
	}
	return comma(s[:length-3-decimal]) + "," + s[length-3-decimal:]
}
