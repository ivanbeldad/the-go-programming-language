// Modify charcount to count letters, digits, and so on in their Unicode categories,
// using functions like unicode.IsLetter.

package main

import (
	"encoding/json"
	"fmt"
	"unicode"
)

func main() {
	count("1.- My string. 2- Another string")
}

func count(s string) {
	m := make(map[string]map[string]int)
	m["letters"] = make(map[string]int)
	m["digits"] = make(map[string]int)
	m["puncts"] = make(map[string]int)
	m["others"] = make(map[string]int)
	for _, r := range s {
		switch {
		case unicode.IsLetter(r):
			m["letters"][string(r)]++
		case unicode.IsDigit(r):
			m["digits"][string(r)]++
		case unicode.IsPunct(r) || unicode.IsSpace(r):
			m["puncts"][string(r)]++
		default:
			m["others"][string(r)]++
		}
	}
	b, _ := json.MarshalIndent(m, "", "  ")
	fmt.Printf("%s\n", string(b))
}
