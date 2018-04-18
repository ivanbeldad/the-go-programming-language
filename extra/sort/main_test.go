package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

const (
	dataDir    = "data/"
	fewFile    = dataDir + "few.txt"
	someFile   = dataDir + "some.txt"
	manyFile   = dataDir + "many.txt"
	fewAmount  = 10
	someAmount = 100
	manyAmount = 3000
)

var few []int
var some []int
var many []int

func init() {
	fewB, fewE := ioutil.ReadFile(fewFile)
	someB, someE := ioutil.ReadFile(someFile)
	manyB, manyE := ioutil.ReadFile(manyFile)
	if fewE != nil || someE != nil || manyE != nil {
		fmt.Printf("Error opening some file.\n")
		os.Exit(1)
	}
	few = parse(fewB)
	some = parse(someB)
	many = parse(manyB)
	fmt.Printf("\n------------------------\n")
	fmt.Printf("Few:\t%5.d\n", fewAmount)
	fmt.Printf("Some:\t%5.d\n", someAmount)
	fmt.Printf("Many:\t%5.d\n", manyAmount)
	fmt.Printf("------------------------\n\n")
}

func parse(b []byte) []int {
	n := make([]int, 0)
	str := string(b)
	str = strings.TrimRight(str, "\r\n")
	strs := strings.Split(str, " ")
	for _, s := range strs {
		val, err := strconv.Atoi(s)
		if err != nil {
			log.Fatal(err)
		}
		n = append(n, val)
	}
	return n
}

func BenchmarkSelectionSortFew(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SelectionSort(append([]int(nil), few...))
	}
}

func BenchmarkDobleSelectionSortFew(b *testing.B) {
	for n := 0; n < b.N; n++ {
		DobleSelectionSort(append([]int(nil), few...))
	}
}

func BenchmarkBubbleSortFew(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BubbleSort(append([]int(nil), few...))
	}
}

func BenchmarkTreeSortFew(b *testing.B) {
	for n := 0; n < b.N; n++ {
		TreeSort(append([]int(nil), few...), true)
	}
}

func BenchmarkSelectionSortSome(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SelectionSort(append([]int(nil), some...))
	}
}

func BenchmarkDobleSelectionSortSome(b *testing.B) {
	for n := 0; n < b.N; n++ {
		DobleSelectionSort(append([]int(nil), some...))
	}
}

func BenchmarkBubbleSortSome(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BubbleSort(append([]int(nil), some...))
	}
}

func BenchmarkTreeSortSome(b *testing.B) {
	for n := 0; n < b.N; n++ {
		TreeSort(append([]int(nil), some...), true)
	}
}

func BenchmarkSelectionSortMany(b *testing.B) {
	for n := 0; n < b.N; n++ {
		SelectionSort(append([]int(nil), many...))
	}
}

func BenchmarkDobleSelectionSortMany(b *testing.B) {
	for n := 0; n < b.N; n++ {
		DobleSelectionSort(append([]int(nil), many...))
	}
}

func BenchmarkBubbleSortMany(b *testing.B) {
	for n := 0; n < b.N; n++ {
		BubbleSort(append([]int(nil), many...))
	}
}

func BenchmarkTreeSortMany(b *testing.B) {
	for n := 0; n < b.N; n++ {
		TreeSort(append([]int(nil), many...), true)
	}
}
