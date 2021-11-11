package main

import "fmt"

func findMinimal(array []int) int {
	var pos int

	for i, n := range array[1:] {
		if array[pos] > n {
			pos = i + 1
		}
	}
	return pos
}

func selectionSort(array *[]int) {
	var i, j int

	for ; i<len((*array)); i++ {
		j = findMinimal((*array)[i:]) + i
		(*array)[i], (*array)[j] = (*array)[j], (*array)[i]
	}
}

func main() {
	var test = []int{8, 5, 2, 9, 5, 6, 3}
	var example = []int{2, 3, 5, 5, 6, 8, 9}

	selectionSort(&test)
	fmt.Println(test, "vs", example)
}
