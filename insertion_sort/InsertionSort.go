package main

import "fmt"

func insertionSort(array *[]int) {
	var j int

	for i:=1; i<len(*array); i++ {
		j = i
		for j > 0 && (*array)[j] < (*array)[j-1] {
			(*array)[j], (*array)[j-1] = (*array)[j-1], (*array)[j]
			j -= 1
		}
	}
}

func main() {
	var array = []int{8, 5, 2, 9, 5, 6, 3}
	var example = []int{2, 3, 5, 5, 6, 8, 9}

	insertionSort(&array)
	fmt.Println(array, "vs", example)
}
