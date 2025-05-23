package main

import "fmt"

func main() {
	list := []int{38, 27, 43, 3, 9, 82, 10}
	final := mergesort(list)
	fmt.Println(final)
}

func mergesort(list []int) []int {
	if len(list) <= 1 {
		return list
	}

	split := len(list)/2 + 1
	left := list[:split]
	right := list[split:]

	sortedleft := mergesort(left)
	sortedright := mergesort(right)

	return merge(sortedleft, sortedright)
}

func merge(left, right []int) []int {
	var val []int
	if left[0] < right[0] {
		left := append(left, right[0])
		val = left
	} else {
		right := append(right, left[0])
		val = right
	}
	return val
}
