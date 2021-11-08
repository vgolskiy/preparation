package main

import "fmt"

func getFibonacciNumber(n int) int {
	var tmp int

	prev := 0
	curr := 1
	if n > 1 {
		for n > 2 {
			tmp = curr
			curr += prev
			prev = tmp
			n -= 1
		}
		return curr
	}
	return prev
}

func main() {
	var num int

	fmt.Print("Please, input number: ")
	fmt.Scanf("%d", &num)
	fmt.Println(getFibonacciNumber(num))
}
