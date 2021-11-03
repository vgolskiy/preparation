package main

import (
	"fmt"
	"math"
)

func sortedSquaredArray(arr *[]int) []int {
	var begin, end, pos int
	n := len(*arr)
	res := make([]int, n)

	end = n - 1
	pos = end

	for begin <= end {
		if math.Abs(float64((*arr)[begin])) < math.Abs(float64((*arr)[end])) {
			res[pos] = int(math.Pow(float64((*arr)[end]), float64(2)))
			end -= 1
		} else {
			res[pos] = int(math.Pow(float64((*arr)[begin]), float64(2)))
			begin += 1
		}
		pos -= 1
	}
	return res
}

func main() {
	fmt.Println(sortedSquaredArray(&[]int{-7, -5, -4, 3, 6, 8, 9}))
}
