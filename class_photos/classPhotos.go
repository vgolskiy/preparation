package main

import (
	"fmt"
	"sort"
)

func classPhotos(reds []int, blues []int) bool {
	var red bool

	sort.Ints(reds)
	sort.Ints(blues)
	n := len(reds)
	if reds[n - 1] < blues[n-1] {
		red = true
	}
	for i:=0; i<n; i++ {
		if red {
			if reds[i] >= blues[i] {
				return false
			}
		} else {
			if reds[i] <= blues[i] {
				return false
			}
		}
	}
	return true
}

func main() {
	var reds = []int{5, 8, 1, 3, 4}
	var blues = []int{6, 9, 2, 4, 5}

	fmt.Println(classPhotos(reds, blues))
}
