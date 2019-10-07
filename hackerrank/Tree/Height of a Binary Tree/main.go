package main

import (
	"fmt"
)

// Node struct
type Node struct {
	value int
	left  *Node
	right *Node
}

func getPaths(arr []*Node, l int) int {
	var nodes []*Node
	for _, node := range arr {
		if node.left != nil {
			nodes = append(nodes, node.left)
		}
		if node.right != nil {
			nodes = append(nodes, node.right)
		}

		if node.left == nil && node.right == nil {
			return l
		}

		return getPaths(nodes, l+1)
	}

	return l
}

func getHeight(node *Node) int {

	return getPaths([]*Node{node}, 0)
}

func main() {
	e := &Node{1, nil, nil}
	e.left = &Node{21, nil, nil}
	e.right = &Node{22, nil, nil}
	e.left.left = &Node{31, nil, nil}
	e.left.left.left = &Node{41, nil, nil}

	fmt.Println(getHeight(e))
}
