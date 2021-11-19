package main

import "fmt"

func getPermutations(arr []int) (res [][]int) {
	var helper func([]int, int)

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n - 1)
				if n % 2 == 1 {
					arr[i], arr[n - 1] = arr[n - 1], arr[i]
				} else {
					arr[0], arr[n - 1] = arr[n - 1], arr[0]
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func main() {
	fmt.Println(getPermutations([]int{1, 2, 3, 4}))
}
