package main

func MinHeightBST(array []int) *BST {
	return buildBST(array, nil, 0, len(array)-1)
}

func buildBST(array []int, bst *BST, startIdx, endIdx int) *BST {
	if endIdx < startIdx {
		return nil
	}

	midIdx := startIdx + (endIdx-startIdx)/2
	value := array[midIdx]
	if bst == nil {
		bst = &BST{Value: value}
	} else {
		bst.Insert(value)
	}
	buildBST(array, bst, startIdx, midIdx-1)
	buildBST(array, bst, midIdx+1, endIdx)
	return bst
}

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) Insert(value int) *BST {
	if value < tree.Value {
		if tree.Left == nil {
			tree.Left = &BST{Value: value}
		} else {
			tree.Left.Insert(value)
		}
	} else {
		if tree.Right == nil {
			tree.Right = &BST{Value: value}
		} else {
			tree.Right.Insert(value)
		}
	}
	return tree
}
