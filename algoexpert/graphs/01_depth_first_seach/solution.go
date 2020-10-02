package main

// Do not edit the class below except
// for the depthFirstSearch method.
// Feel free to add new properties
// and methods to the class.
type Node struct {
	Name     string
	Children []*Node
}

type NodeStack []*Node

func (s NodeStack) Len() int {
	return len(s)
}

func (s *NodeStack) Push(node *Node) {
	*s = append(*s, node)
}

func (s *NodeStack) Pop() *Node {
	old := *s
	n := len(old)
	node := old[n-1]
	*s = old[:n-1]
	return node
}

func (n *Node) DepthFirstSearch(array []string) []string {
	result := make([]string, 0)
	if n == nil {
		return result
	}

	stack := make(NodeStack, 0)
	stack.Push(n)
	for len(stack) != 0 {
		node := stack.Pop()
		result = append(result, node.Name)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack.Push(node.Children[i])
		}
	}

	return result
}
