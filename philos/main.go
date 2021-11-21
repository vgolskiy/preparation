package main

import (
	"fmt"
	"os"
)

func main() {
	var argc int
	argc = len(os.Args[1:])
	fmt.Println(argc)
	args := os.Args[1:]
	if argc != 4 {
		error(1)
	} else {
		fmt.Println(args)
	}
}