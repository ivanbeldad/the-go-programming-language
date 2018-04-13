// Write a general-purpose unit-conversion program analogous to cf that reads
// numbers from its command-line arguments or from the standard input if there are no arguments,
// and converts each number into units like temperature in Celsius and Fahrenheit,
// length in feet and meters, weight in pounds and kilograms, and the like.

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Kilogram ...
type Kilogram float64

// Pound ...
type Pound float64

const (
	poundRatio = 0.45359237
)

func main() {
	var input string
	if len(os.Args) == 1 {
		input = inputFromStd()
	} else {
		input = os.Args[1]
	}
	weigth, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	poundsInKg := KtoP(Kilogram(weigth))
	kgInPounds := PtoK(Pound(weigth))

	fmt.Printf("%s are %s\n", Kilogram(weigth), poundsInKg)
	fmt.Printf("%s are %s\n", Pound(weigth), kgInPounds)
}

func inputFromStd() string {
	reader := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	line := scanner.Text()
	return line
}

func (p Pound) String() string    { return fmt.Sprintf("%g lbs", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%g kg", k) }

// PtoK ...
func PtoK(p Pound) Kilogram { return Kilogram(p * poundRatio) }

// KtoP ...
func KtoP(k Kilogram) Pound { return Pound(k / poundRatio) }
