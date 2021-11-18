package main

import "fmt"

func generator(s string, open, close, n int) {
	if len(s) == 2*n {
		fmt.Println(s)
		return
	}
	if open < n {
		generator(s + "(", open+1, close, n)
	}
	if close < open {
		generator(s + ")", open, close+1, n)
	}
}

func main() {
	generator("", 0, 0, 3)
}
