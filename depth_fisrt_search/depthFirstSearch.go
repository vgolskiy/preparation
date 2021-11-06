package main

import "fmt"

type Node struct {
	Children	[]Node
	Name		string
}

func (n *Node) addChild(name string) *Node {
	n.Children = append(n.Children, Node{Children: nil, Name: name})
	return n
}

func (n Node) depthFirstSearch(array *[]string) []string {
	*array = append(*array, n.Name)
	for _, child := range n.Children {
		child.depthFirstSearch(array)
	}
	return *array
}

func main() {
	graph := Node{Children: nil, Name: "A"}
	graph.addChild("B").addChild("C").addChild("D")
	graph.Children[0].addChild("E").addChild("F")
	graph.Children[2].addChild("G").addChild("H")
	graph.Children[0].Children[1].addChild("I").addChild("J")
	graph.Children[2].Children[0].addChild("K")

	var array []string
	var ex = []string{"A", "B", "E", "F", "I", "J", "C", "D", "G", "K", "H"}
	fmt.Println(graph.depthFirstSearch(&array), "vs", ex)
}
