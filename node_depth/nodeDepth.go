package main

import (
	"fmt"
	"io"
	"os"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

type BST struct {
	Root *Node
}

func calcDepthSumHelper(node *Node, sums *[]int, lvl int) {
	if node == nil {
		return
	}
	*sums = append(*sums, lvl)
	calcDepthSumHelper(node.Left, sums, lvl+1)
	calcDepthSumHelper(node.Right, sums, lvl+1)
}

func calcDepthSum(tree *BST) int {
	var res int
	var sums []int

	calcDepthSumHelper(tree.Root, &sums, 0)
	for _, num := range sums {
		res += num
	}
	return res
}

func (node *Node) insert(value int) {
	if node == nil {
		return
	} else if value < node.Value {
		if node.Left == nil {
			node.Left = &Node{Value: value, Left: nil, Right: nil}
		} else {
			node.Left.insert(value)
		}
	} else {
		if node.Right == nil {
			node.Right = &Node{Value: value, Left: nil, Right: nil}
		} else {
			node.Right.insert(value)
		}
	}
}

func (tree *BST) insert(value int) *BST {
	if tree.Root == nil {
		tree.Root = &Node{Value: value, Left: nil, Right: nil}
	} else {
		tree.Root.insert(value)
	}
	return tree
}

func print(w io.Writer, node *Node, ns int, ch rune) {
	if node == nil {
		return
	}
	for i := 0; i < ns; i++ {
		fmt.Fprint(w, " ")
	}
	fmt.Fprintf(w, "%c:%v\n", ch, node.Value)
	print(w, node.Left, ns+2, 'L')
	print(w, node.Right, ns+2, 'R')
}

func main() {
	tree := &BST{}
	tree.insert(5).
		insert(4).
		insert(3).
		insert(2).
		insert(1).
		insert(6).
		insert(7).
		insert(8).
		insert(9).
		insert(10)
	print(os.Stdout, tree.Root, 0, 'M')
	summa := 25
	fmt.Println(summa, "vs", calcDepthSum(tree))
}
