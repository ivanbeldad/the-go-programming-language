// Modify reverse to reverse the characters of a []byte slice that represents a
// UTF-8-encoded string, in place. Can you do it without allocating new memory?

package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	b := []byte("Hello 本 World!")
	fmt.Printf("Original: \t\t%q\n", b)
	reverse(&b)
	fmt.Printf("Reversed (standard): \t%q\n", b)
	reverse(&b)
	reverseBuffer(&b)
	fmt.Printf("Reversed (buffer): \t%q\n", b)
	reverseBuffer(&b)
	reverseBuilder(&b)
	fmt.Printf("Reversed (builder): \t%q\n", b)
	reverseBuilder(&b)
	reverseCopy(&b)
	fmt.Printf("Reversed (copy): \t%q\n", b)
	reverseCopy(&b)
}

func reverse(b *[]byte) {
	runes := []rune(string(*b))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	*b = []byte(string(runes))
}

func reverseBuffer(b *[]byte) {
	var buffer bytes.Buffer
	runes := []rune(string(*b))
	for i := len(runes) - 1; i >= 0; i-- {
		buffer.WriteRune(runes[i])
	}
	*b = buffer.Bytes()
}

func reverseBuilder(b *[]byte) {
	runes := []rune(string(*b))
	var builder strings.Builder
	for i := len(runes) - 1; i >= 0; i-- {
		builder.WriteRune(runes[i])
	}
	*b = []byte(builder.String())
}

func reverseCopy(b *[]byte) {
	runes := []rune(string(*b))
	res := make([]rune, len(runes))
	for i, j := len(runes)-1, 0; i >= 0; i, j = i-1, j+1 {
		copy(res[j:], runes[i:i+1])
	}
	*b = []byte(string(res))
}
