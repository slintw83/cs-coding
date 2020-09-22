package main

import "fmt"

// Do not edit the class below except for
// the insert, contains, and remove methods.
// Feel free to add new properties and methods
// to the class.
type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	fmt.Println("Inserting ", value)
	return insert(tree, value)
}

func insert(tree *BST, value int) *BST {
	if tree == nil {
		return &BST{Value: value}
	}

	if value < tree.Value {
		tree.Left = insert(tree.Left, value)
	} else if value >= tree.Value {
		tree.Right = insert(tree.Right, value)
	}

	return tree
}

func (tree *BST) Contains(value int) bool {
	fmt.Println("Contains ", value)
	node := tree
	for node != nil {
		if node.Value == value {
			return true
		} else if value < node.Value {
			node = node.Left
		} else if value > node.Value {
			node = node.Right
		}
	}
	return false
}

func (tree *BST) Remove(value int) *BST {
	fmt.Println("Remove ", value)

	if tree.Left == nil && tree.Right == nil {
		return tree
	}

	tmp := remove(tree, value)
	tree.Value = tmp.Value
	tree.Right = tmp.Right
	tree.Left = tmp.Left
	return tree
}

func remove(tree *BST, value int) *BST {
	if tree == nil {
		return nil
	}

	if value < tree.Value {
		tree.Left = remove(tree.Left, value)
	} else if value > tree.Value {
		tree.Right = remove(tree.Right, value)
	} else {
		if tree.Right == nil {
			return tree.Left
		}
		if tree.Left == nil {
			return tree.Right
		}

		tmp := tree
		tree = min(tmp.Right)
		tree.Right = deleteMin(tmp.Right)
		tree.Left = tmp.Left
	}

	return tree
}

func min(tree *BST) *BST {
	if tree.Left == nil {
		return tree
	}
	return min(tree.Left)
}

func deleteMin(tree *BST) *BST {
	if tree.Left == nil {
		return tree.Right
	}

	tree.Left = deleteMin(tree.Left)
	return tree
}
