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

type Stack struct {
	Node  *Node
	Depth int
}

func calcDepthSum(tree *BST) int {
	var res int
	var stack = []Stack{
		{
			Node:  tree.Root,
			Depth: 0,
		},
	}
	for len(stack) > 0 {
		node, depth, tmp := stack[len(stack)-1].Node, stack[len(stack)-1].Depth, stack[:len(stack)-1]
		stack = tmp
		tmp = nil
		if node == nil {
			continue
		}
		res += depth
		stack = append(stack, Stack{Node: node.Left, Depth: depth + 1})
		stack = append(stack, Stack{Node: node.Right, Depth: depth + 1})
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
