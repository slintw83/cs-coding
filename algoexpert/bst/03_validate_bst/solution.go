package main

import "math"

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Validate() bool {
	return validate(tree, math.MinInt32, math.MaxInt32)
}

func validate(tree *BST, minValue, maxValue int) bool {
	if tree.Value < minValue || tree.Value >= maxValue {
		return false
	}

	if tree.Left != nil && !validate(tree.Left, minValue, tree.Value) {
		return false
	}

	if tree.Right != nil && !validate(tree.Right, tree.Value, maxValue) {
		return false
	}

	return true
}
