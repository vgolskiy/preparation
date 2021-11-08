package main

import "fmt"

func binarySearchHelper(array []int, target int, left int, right int) int {
	if left > right {
		return -1
	}
	pos := (left + right) / 2
	if array[pos] == target {
		return pos
	} else if array[pos] > target {
		return binarySearchHelper(array, target, left, pos - 1)
	}
	return binarySearchHelper(array, target, pos + 1, right)
}

func binarySearch(array []int, target int) int {
	var left, right int

	right = len(array) - 1
	return binarySearchHelper(array, target, left, right)
}

func main() {
	var array = []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}

	fmt.Println(binarySearch(array, 33), "vs", 3)
}
