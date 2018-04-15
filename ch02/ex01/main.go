// Add types, constants, and functions to tempconv for processing temperatures in the Kelvin
// scale, where zero Kelvin is −273.15°C and a difference of 1K has the same magnitude as 1°C.

package main

import "fmt"

// Celsius ...
type Celsius float64

// Fahrenheit ...
type Fahrenheit float64

// Kelvin ...
type Kelvin float64

const (
	// AbsoluteZeroC absolute temperature in Celsius
	AbsoluteZeroC Celsius = -273.15
	// FreezingC freezing temperature in Celsius
	FreezingC Celsius = 0
	// BoilingC boiling temperature in Celsius
	BoilingC Celsius = 100
)

func main() {
	var temp Celsius
	fmt.Printf("%s", CToK(temp))
}

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }

// CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KToC converts a Kelvin temperature to Celsius
func KToC(k Kelvin) Celsius { return Celsius(k + 273.15) }

// CToK converts a Celsius temperature to Kelvin
func CToK(c Celsius) Kelvin { return Kelvin(c - 273.15) }

// KToF converts a Kelvin temperature to Farenheit
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }
