// Write a variadic version of strings.Join.

package main

import (
	"fmt"
	"strings"
)

func main() {
	strings := []string{
		"first",
		"second",
		"third",
		"fourth",
		"fifth",
	}
	fmt.Println(join("---", strings...))
}

func join(j string, strs ...string) (s string) {
	if len(strs) == 0 {
		return ""
	}
	buf := strings.Builder{}
	for _, str := range strs {
		buf.WriteString(str)
		buf.WriteString(j)
	}
	return buf.String()[:len(buf.String())-len(j)]
}
