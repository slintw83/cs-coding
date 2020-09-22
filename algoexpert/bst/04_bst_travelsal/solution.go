package main

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func (tree *BST) InOrderTraverse(array []int) []int {
	array = inorder(tree, array)
	return array
}

func inorder(tree *BST, result []int) []int {
	if tree == nil {
		return result
	}

	result = inorder(tree.Left, result)
	result = append(result, tree.Value)
	result = inorder(tree.Right, result)

	return result
}

func (tree *BST) PreOrderTraverse(array []int) []int {
	array = preorder(tree, array)
	return array
}

func preorder(tree *BST, result []int) []int {
	if tree == nil {
		return result
	}

	result = append(result, tree.Value)
	result = preorder(tree.Left, result)
	result = preorder(tree.Right, result)
	return result
}

func (tree *BST) PostOrderTraverse(array []int) []int {
	array = postorder(tree, array)
	return array
}

func postorder(tree *BST, result []int) []int {
	if tree == nil {
		return result
	}

	result = postorder(tree.Left, result)
	result = postorder(tree.Right, result)
	result = append(result, tree.Value)
	return result
}
