package main

import (
	"fmt"
	"sort"
)

func threeNumSum(arr []int, target int) (res [][3]int) {
	var left, right, currSum int
	sort.Ints(arr)
	n := len(arr)
	for i:=0; i < n-2; i++ {
		left = i + 1
		right = n - 1
		for left < right {
			currSum = arr[i] + arr[left] + arr[right]
			if currSum == target {
				res = append(res, [3]int{arr[i], arr[left], arr[right]})
				left += 1
				right -= 1
			} else if currSum > target {
				right -= 1
			} else {
				left += 1
			}
		}
	}
	return
}

func main() {
	example := [][3]int{{-8, 2, 6}, {-8, 3, 5}, {-6, 1, 5}}
	fmt.Println(threeNumSum([]int{12, 3, 1, 2, -6, 5, -8, 6}, 0), "vs", example)
}
