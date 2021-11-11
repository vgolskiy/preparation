package main

import "fmt"

func firstNonRepeat(s string) int {
	alp := make(map[rune]int)

	for _, ch := range s {
		if _, ok := alp[ch]; ok {
			alp[ch] += 1
		} else {
			alp[ch] = 1
		}
	}
	for i, ch := range s {
		if alp[ch] == 1 {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(1, "vs", firstNonRepeat("abcdcaf"))
}
