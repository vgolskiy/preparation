package main

import "fmt"

func bubbleSort(array *[]int) {
	var qty int

	more := true
	for more {
		more = false
		for i:=0; i<(len(*array)-1-qty); i++ {
			if (*array)[i] > (*array)[i+1] {
				(*array)[i], (*array)[i+1] = (*array)[i+1], (*array)[i]
				more = true
			}
		}
		qty += 1
	}
}

func main() {
	var array = []int{8, 5, 2, 9, 5, 6, 3}
	var example = []int{2, 3, 5, 5, 6, 8, 9}

	bubbleSort(&array)
	fmt.Println(array, "vs", example)
}
