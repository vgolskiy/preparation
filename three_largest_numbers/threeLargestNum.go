package main

import "fmt"

func fillThree(three *[]int, num int) {
	*three = append(*three, num)
	n := len(*three)
	if n == 1 {
		return
	}
	if n == 3 && (*three)[1] > (*three)[2] {
		(*three)[1], (*three)[2] = (*three)[2], (*three)[1]
	}
	if (*three)[0] > (*three)[1] {
		(*three)[0], (*three)[1] = (*three)[1], (*three)[0]
	}
}

func setThreeNum(three *[]int, num int) {
	if (*three)[2] < num {
		(*three)[0], (*three)[1], (*three)[2] = (*three)[1], (*three)[2], num
	} else if (*three)[1] < num {
		(*three)[0], (*three)[1], (*three)[2] = (*three)[0], num, (*three)[2]
	} else {
		(*three)[0], (*three)[1], (*three)[2] = num, (*three)[1], (*three)[2]
	}
}

func threeLargestNums(array []int) []int {
	var three []int

	for i, num := range array {
		if i < 3 {
			fillThree(&three, num)
		} else {
			if three[0] < num {
				setThreeNum(&three, num)
			}
		}
	}
	return three
}

func main() {
	var array = []int{141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7}
	var three = []int{18, 141, 541}

	fmt.Println(threeLargestNums(array), "vs", three)
}
