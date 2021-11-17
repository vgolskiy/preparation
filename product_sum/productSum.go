package main

import "fmt"

func productSum(array []interface{}, depth int) int {
	var res int

	for _, el := range array {
		if v, ok := el.(int); ok {
			res += v
		} else {
			res += productSum(el.([]interface{}), depth + 1)
		}
	}
	return res * depth
}

func main() {
	var array = []interface{}{5, 2, []interface{}{7, -1}, 3, []interface{}{6, []interface{}{-13, 8}, 4}}
	fmt.Println(12, "vs", productSum(array, 1))
}
