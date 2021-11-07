package main

import (
	"fmt"
	"math"
	"sort"
)

func tandemBycicle(reds []int, blues []int, fastest bool) int {
	var res int

	sort.Ints(reds)
	sort.Ints(blues)
	n := len(reds)
	for i:=0; i<n; i++ {
		if fastest {
			res += int(math.Max(float64(reds[n - (i + 1)]), float64(blues[i])))
		} else {
			res += int(math.Max(float64(reds[i]), float64(blues[i])))
		}
	}
	return res
}

func main() {
	var reds = []int{5, 5, 3, 9, 2}
	var blues = []int{3, 6, 7, 2, 1}

	fastest := false
	expected := 25
	if fastest {
		expected = 32
	}
	fmt.Println(expected, "vs", tandemBycicle(reds, blues, fastest))
}