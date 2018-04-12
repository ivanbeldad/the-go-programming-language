// Modify the echo program to print the index and value of each of its arguments, one per line.
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var result string
	for index, arg := range os.Args {
		result += strconv.Itoa(index) + " " + arg + "\n"
	}
	result = result[0 : len(result)-1]
	fmt.Println(result)
}
