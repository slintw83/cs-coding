package main

import "fmt"

// Do not edit the class below except
// for the breadthFirstSearch method.
// Feel free to add new properties
// and methods to the class.
type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) BreadthFirstSearch(array []string) []string {
	if n == nil {
		return []string{}
	}

	result := make([]string, 0)
	q := make([]*Node, 0)
	q = append(q, n)
	for len(q) != 0 {
		node := q[0]
		q = q[1:]
		fmt.Print(node.Name, " ")
		result = append(result, node.Name)
		for _, child := range node.Children {
			q = append(q, child)
		}
	}
	return result
}
