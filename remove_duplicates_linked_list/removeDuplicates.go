package main

import "fmt"

type List struct {
	Value	int
	Next	*List
}

func (l *List) addMany(values []int) *List {
	current := l
	for current.Next != nil {
		current = current.Next
	}
	for _, v := range values {
		current.Next = &List{Value: v, Next: nil}
		current = current.Next
	}
	return l
}

func (l *List) getNodesArray() []int {
	var res []int

	current := l
	for current != nil {
		res = append(res, current.Value)
		current = current.Next
	}
	return res
}

func removeDuplicates(l *List) {
	current := l
	for current != nil {
		next := current.Next
		for next != nil && next.Value == current.Value {
			next = next.Next
		}
		current.Next = next
		current = current.Next
	}
}

func main() {
	test := List{Value: 1, Next: nil}
	test.addMany([]int{1, 3, 4, 4, 4, 5, 6, 6})
	fmt.Println(test.getNodesArray())
	expected := List{Value: 1, Next: nil}
	expected.addMany([]int{3, 4, 5, 6})
	fmt.Println(expected.getNodesArray())
	removeDuplicates(&test)
	fmt.Println(test.getNodesArray(), "vs", expected.getNodesArray())
}