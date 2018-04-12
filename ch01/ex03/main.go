// Experiment to measure the difference in running time between our potentially inefficient
// versions and the one that uses strings.Join.

package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const size = 100000
const word = "Hello"

func main() {
	var start int64

	start = time.Now().UnixNano()
	concat(size, word)
	concatTime := time.Now().UnixNano() - start

	start = time.Now().UnixNano()
	join(size, word)
	joinTime := time.Now().UnixNano() - start

	strConcatTime := strconv.Itoa(int(concatTime))
	strJoinTime := strconv.Itoa(int(joinTime))

	fmt.Println("Concat time : " + strConcatTime + " µs")
	fmt.Println("Join time   : " + strJoinTime + " µs")

	timesFaster := int(concatTime / joinTime)

	fmt.Println("Join is " + strconv.Itoa(timesFaster) + " times faster than concatenation")
}

func concat(size int, word string) string {
	text := ""
	for i := 0; i < size; i++ {
		text += word + " "
	}
	return text
}

func join(size int, word string) string {
	words := make([]string, size)
	for i := 0; i < size; i++ {
		words[i] = word
	}
	text := strings.Join(words, " ")
	return text
}
