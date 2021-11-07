package main

import (
	"fmt"
	"sort"
)

func minWaitingTime(queue []int) int {
	var res, n int

	n = len(queue)
	sort.Ints(queue)
	for i, q := range queue {
		res += q * (n - (i + 1))
	}
	return res
}

func main() {
	var queue = []int{3, 2, 1, 2, 6}
	fmt.Println(17, "vs", minWaitingTime(queue))
}
