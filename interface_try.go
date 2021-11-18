package main

import "fmt"

type intIncrement int

// no pointer - no incremention (1 always)
func (inc *intIncrement) Increment() int {
	*inc++
	return int(*inc)
}

func main() {
	myInt := intIncrement(0)

	for i:=0; i < 10; i++ {
		fmt.Println(myInt.Increment())
	}
	fmt.Println()
}
