package main

import "fmt"

func twoNumSum(arr *[]int, target int) []int {
	buff := make(map[int]int)

	for _, num := range *arr {
		if value, found := buff[num]; found {
			return []int{value, num}
		}
		buff[target-num] = num
	}
	return []int{}
}

func main() {
	fmt.Println(twoNumSum(&[]int{2, 9, -1, 3, 10, 11}, 8))
}
