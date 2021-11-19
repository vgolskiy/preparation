package main

import "fmt"

func verifyPeak(arr []int, i int) bool {
	return arr[i] > arr[i-1] && arr[i] > arr[i+1]
}

func longestPeak(arr []int) int {
	var left, right int

	i := 1
	max_len := 0
	for i < len(arr) - 1 {
		if verifyPeak(arr, i) {
			left = i - 2
			for left >= 0 && arr[left] < arr[left+1] {
				left--
			}
			right = i + 2
			for right < len(arr) && arr[right] < arr[right-1] {
				right++
			}
			if max_len < right - left - 1 {
				max_len = right - left - 1
			}
			i = right
		} else {
			i++
		}
	}
	return max_len
}

func main() {
	arr := []int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3}
	expected := 6
	fmt.Println(longestPeak(arr), "vs", expected)
}
