package main

import (
	"fmt"
	"os"
)

func main() {
	s := []rune(os.Args[1])
	for i, r := range s {
		if i > 0 {
			r -= int32(i)
			s[i] = r
		}
	}
	fmt.Println(string(s))
}
