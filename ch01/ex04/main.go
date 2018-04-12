// Modify dup2 to print the names of all files in which each duplicated line occurs.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	filesCount := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, filesCount)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, filesCount)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, filesCount[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int, filesCount map[string]string) {
	added := make(map[string]bool)
	input := bufio.NewScanner(f)
	for input.Scan() {
		content := input.Text()
		counts[content]++
		if !added[content] {
			filesCount[content] += f.Name() + " "
		}
		added[content] = true
	}
	// NOTE: ignoring potential errors from input.Err()
}
