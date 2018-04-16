// Write a program wordfreq to report the frequency of each word in an input text file.
// Call input.Split(bufio.ScanWords) before the first call to Scan to break the input
// into words instead of lines.

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	freq := wordfreq("hello world hello worldsito")
	freqText, _ := json.MarshalIndent(freq, "", "  ")
	fmt.Printf("%s", freqText)
}

func wordfreq(s string) map[string]int {
	reader := strings.NewReader(s)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	words := make(map[string]int)
	for scanner.Scan() {
		words[scanner.Text()]++
	}
	return words
}
