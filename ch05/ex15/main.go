// Write variadic functions max and min, analogous to sum. What should these
// functions do when called with no arguments? Write variants that require at least one argument.

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%d\n", max(1, 3, 7, 2))
	fmt.Printf("%d\n", min(-1, 3, 7, 2))
	fmt.Printf("%d\n", max(6))
	fmt.Printf("%d\n", maxForce(43))
	fmt.Printf("%d\n", maxForce(43, -32, 1000))
	fmt.Printf("%d\n", min())
}

func max(nums ...int) (max int) {
	if len(nums) == 0 {
		panic("max requires at least one int")
	}
	max = nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return
}

func min(nums ...int) (min int) {
	if len(nums) == 0 {
		panic("min requires at least one int")
	}
	min = nums[0]
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return
}

func maxForce(n int, nums ...int) (max int) {
	max = n
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return
}

func minForce(n int, nums ...int) (min int) {
	min = n
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return
}
