package main

import "fmt"

func fisrt_non_repeat(s *string) int {
	buf := make(map[rune]int)

	for _, ch := range *s {
		if _, found := buf[ch]; found {
			buf[ch] += 1
		} else {
			buf[ch] = 1
		}
	}
	for i, ch := range *s {
		if buf[ch] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	s := "abracadabra"
	fmt.Println(fisrt_non_repeat(&s))
}
