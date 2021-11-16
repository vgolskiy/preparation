package main

import "fmt"

func threeNumSort(arr *[]int, order *[3]int) {
	n := len(*arr)
	firstIdx := 0
	for i:=0; i<n; i++ {
		if (*arr)[i] == (*order)[0] {
			(*arr)[i], (*arr)[firstIdx] = (*arr)[firstIdx], (*arr)[i]
			firstIdx += 1
		}
	}
	lastIdx := n - 1
	for i:=n-1; i>firstIdx-1; i-- {
		if (*arr)[i] == (*order)[2] {
			(*arr)[i], (*arr)[lastIdx] = (*arr)[lastIdx], (*arr)[i]
			lastIdx -= 1
		}
	}
}

func main() {
	array := []int{1, 0, 0, -1, -1, 0, 1, 1}
	order := [3]int{0, 1, -1}
	expected := []int{0, 0, 0, 1, 1, 1, -1, -1}
	threeNumSort(&array, &order)
	fmt.Println(array, "vs", expected)
}
