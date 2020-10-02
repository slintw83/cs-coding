package main

type AncestralTree struct {
	Name     string
	Ancestor *AncestralTree
}

func GetYoungestCommonAncestor(top, descendantOne, descendantTwo *AncestralTree) *AncestralTree {
	h1 := getHeight(top, descendantOne)
	h2 := getHeight(top, descendantTwo)
	diff := Abs(h1 - h2)
	var low, high *AncestralTree
	if h1 > h2 {
		low = descendantTwo
		high = descendantOne
	} else {
		low = descendantOne
		high = descendantTwo
	}

	for i := 0; i < diff; i++ {
		high = high.Ancestor
	}

	for low != high {
		low = low.Ancestor
		high = high.Ancestor
	}

	return low
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getHeight(top, node *AncestralTree) int {
	h := 0
	for node != top {
		h++
		node = node.Ancestor
	}
	return h
}
