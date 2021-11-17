package main

import (
	"fmt"
	"math"
	"sort"
)

func smallestDifference(arr1 []int, arr2 []int) [2]int {
	var i, j int
	var res [2]int

	sort.Ints(arr1)
	sort.Ints(arr2)
	diff := math.Inf(1)
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] == arr2[j] {
			return [2]int{arr1[i], arr2[j]}
		} else if arr1[i] < arr2[j] {
			if diff > float64(arr2[j] - arr1[i]) {
				diff = float64(arr2[j] - arr1[i])
				res = [2]int{arr1[i], arr2[j]}
			}
			i++
		} else {
			if diff > float64(arr1[i] - arr2[j]) {
				diff = float64(arr1[i] - arr2[j])
				res = [2]int{arr1[i], arr2[j]}
			}
			j++
		}
	}
	return res
}

func main() {
	arr1 := []int{-1, 5, 10, 20, 28, 3}
	arr2 := []int{26, 134, 135, 15, 17}
	fmt.Println(smallestDifference(arr1, arr2), "vs", [2]int{28, 26})
}
