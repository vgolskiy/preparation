package main

import "fmt"

func generateDocument(chars string, doc string) bool {
	if len(doc) > 0 {
		qty := make(map[rune]int)

		for _, ch := range chars {
			if _, ok := qty[ch]; ok {
				qty[ch] += 1
			} else {
				qty[ch] = 1
			}
		}
		for _, ch := range doc {
			if v, ok := qty[ch]; ok {
				if v-1 >= 0 {
					qty[ch] -= 1
				} else {
					return false
				}
			} else {
				return false
			}
		}
	}
	return true
}

func main() {
	characters := "Bste!hetsi ogEAxpelrt x "
	document := "AlgoExpert is the Best!"
	fmt.Println(true, "vs", generateDocument(characters, document))
}
