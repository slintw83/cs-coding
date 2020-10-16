package main

func GetPermutations(array []int) [][]int {
	result := make([][]int, 0)
	permute(0, array, &result)
	return result
}

func permute(start int, array []int, result *[][]int) {
	if start == len(array)-1 {
		*result = append(*result, clone(array))
		return
	}

	for j := start; j < len(array); j++ {
		swap(array, start, j)
		permute(start+1, array, result)
		swap(array, start, j)
	}
}

func swap(array []int, i, j int) {
	array[i], array[j] = array[j], array[i]
}

func clone(array []int) []int {
	newArray := make([]int, len(array))
	copy(newArray, array)
	return newArray
}
