// Use panic and recover to write a function that contains no return statement
// yet returns a non-zero value.

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Result: %d\n", noReturn())
}

func noReturn() (n int) {
	n = 10
	defer func() {
		recover()
	}()
	panic(n)
}
