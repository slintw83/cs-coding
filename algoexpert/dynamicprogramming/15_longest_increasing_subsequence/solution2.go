package main

func LongestIncreasingSubsequence(input []int) []int {
	n := len(input)
	sequences := make([]int, n)
	indices := make([]int, n+1)
	for i := range input {
		sequences[i] = -1
		indices[i] = -1
	}
	length := 0
	for i, num := range input {
		newLength := binarySearch(1, length, indices, input, num)
		sequences[i] = indices[newLength-1]
		indices[newLength] = i
		length = Max(length, newLength)
	}

	result := make([]int, 0)
	i := indices[length]
	for i != -1 {
		result = append(result, input[i])
		i = sequences[i]
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func binarySearch(start int, end int, indices []int, input []int, num int) int {
	if start > end {
		return start
	}
	middle := start + (end-start)/2
	if input[indices[middle]] < num {
		start = middle + 1
	} else {
		end = middle - 1
	}
	return binarySearch(start, end, indices, input, num)
}
