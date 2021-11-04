package main

import (
	"fmt"
	"sort"
)

func calculateChange(coins *[]int) int {
	var max_change int = 0

	if len(*coins) > 0 {
		sort.Ints(*coins)
		for _, coin := range *coins {
			if max_change+1 < coin {
				break
			}
			max_change += coin
		}
	}
	return max_change + 1
}

//var coins []int = []int{1, 1, 1, 1, 1, 1}
var coins []int = []int{5, 7, 1, 1, 2, 3, 22}

func main() {
	fmt.Println(calculateChange(&coins))
}
