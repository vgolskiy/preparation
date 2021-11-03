package main

import "fmt"

func valSubseq(s *[]int, ss *[]int) bool {
	var i, j, ns, nss int

	i = 0
	j = 0
	ns = len(*s)
	nss = len(*ss)
	for i < ns && j < nss {
		if (*s)[i] == (*ss)[j] {
			j += 1
		}
		i += 1
	}
	return j == nss
}

func main() {
	fmt.Println(valSubseq(&[]int{5, 1, 22, 25, 6, -1, 8, 10}, &[]int{1, 6, -1, 10}))
}
