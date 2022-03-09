package main

import "fmt"

func find_sum_memb(arr *[]int, s *int) []int {
	tmp := make(map[int]int)
	for _, n := range *arr {
		if value, found := tmp[n]; found {
			return []int{value, n}
		} else {
			tmp[*s-n] = n
		}
	}
	return []int{}
}

func main() {
	arr := []int{1, 2, 3, 4, -5}
	s := -1
	fmt.Println(find_sum_memb(&arr, &s))
}
