package main

import (
	"fmt"
)

type tree struct {
	Value       int
	Left, Right *tree
}

func main() {
	// v := []int{1, 3, 3, 2}
	v := []int{5, 9, 1, 10, 7, 3, 6, 8, 4}
	fmt.Printf("%v\n", TreeSort(append([]int(nil), v...), true))
	fmt.Printf("%v\n", SelectionSort(append([]int(nil), v...)))
	fmt.Printf("%v\n", BubbleSort(append([]int(nil), v...)))
	fmt.Printf("%v\n", DobleSelectionSort(append([]int(nil), v...)))
}

// TreeSort ...
func TreeSort(n []int, ascend bool) []int {
	var t *tree
	for _, val := range n {
		t = add(t, val)
	}
	sort := make([]int, 0)
	if ascend {
		sort = *extractMin(t, &sort)
	} else {
		sort = *extractMax(t, &sort)
	}
	return sort
}

func add(t *tree, n int) *tree {
	if t == nil {
		return &tree{Value: n}
	}
	if n > t.Value {
		t.Right = add(t.Right, n)
		return t
	}
	t.Left = add(t.Left, n)
	return t
}

func extractMin(t *tree, nums *[]int) *[]int {
	if t.Left != nil {
		nums = extractMin(t.Left, nums)
	}
	*nums = append(*nums, t.Value)
	if t.Right != nil {
		nums = extractMin(t.Right, nums)
	}
	return nums
}

func extractMax(t *tree, nums *[]int) *[]int {
	if t.Right != nil {
		nums = extractMax(t.Right, nums)
	}
	*nums = append(*nums, t.Value)
	if t.Left != nil {
		nums = extractMax(t.Left, nums)
	}
	return nums
}

// SelectionSort ...
func SelectionSort(n []int) []int {
	for i := 0; i < len(n)-1; i++ {
		min := i
		j := i + 1
		for ; j < len(n); j++ {
			if n[min] > n[j] {
				min = j
			}
		}
		if min != i {
			n[i], n[min] = n[min], n[i]
		}
	}
	return n
}

// BubbleSort ...
func BubbleSort(n []int) []int {
	for i := len(n) - 1; i >= 1; i-- {
		for p1, p2 := 0, 1; p2 <= i; p1, p2 = p1+1, p2+1 {
			if n[p1] > n[p2] {
				n[p1], n[p2] = n[p2], n[p1]
			}
		}
	}
	return n
}

// DobleSelectionSort ...
func DobleSelectionSort(n []int) []int {
	min := 0
	max := 0
	for i := 0; i < len(n)-1; i++ {
		if n[i] < n[min] {
			min = i
		}
		n[i], n[min] = n[min], n[i]
	}
	for i := 0; i < len(n)-1; i++ {
		if n[i] > n[max] {
			max = i
		}
		n[i], n[max] = n[max], n[i]
	}
	for fi, li := 0, len(n)-1; fi < li; fi, li = fi+1, li-1 {
		minI := fi
		maxI := li
		for i := fi + 1; i < li; i++ {
			if n[i] < n[minI] {
				minI = i
			} else if n[i] > n[maxI] {
				maxI = i
			}
		}
		n[fi], n[minI] = n[minI], n[fi]
		n[li], n[maxI] = n[maxI], n[li]
	}
	return n
}
