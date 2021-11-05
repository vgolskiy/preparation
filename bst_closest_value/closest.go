package main

import (
	"fmt"
	"io"
	"math"
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

func findClosest(node *Node, value int, closest int) int {
	if node == nil {
		return closest
	}
	if node.Value == value {
		return node.Value
	}
	if math.Abs(float64(node.Value-value)) < math.Abs(float64(closest-value)) {
		closest = node.Value
	}
	if node.Value < value {
		return findClosest(node.Right, value, closest)
	} else {
		return findClosest(node.Left, value, closest)
	}
}

func (tree *BST) findClosestBSTValue(value int) int {
	return findClosest(tree.Root, value, tree.Root.Value)
}

func main() {
	tree := &BST{}
	tree.insert(100).
		insert(-20).
		insert(-50).
		insert(-15).
		insert(-60).
		insert(50).
		insert(60).
		insert(55).
		insert(85).
		insert(15).
		insert(5).
		insert(-10)
	print(os.Stdout, tree.Root, 0, 'M')
	to_find := 35
	fmt.Printf("Closest to %d is %d\n", to_find, tree.findClosestBSTValue(to_find))
}
