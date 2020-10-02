package main

import "fmt"

// LongestIncreasingSubsequence O(N) space and O(N^2) time
func LongestIncreasingSubsequence(input []int) []int {
	sequence := make([]int, len(input))
	length := make([]int, len(input))

	for i := range input {
		sequence[i] = -1
		length[i] = 1
	}

	for i := 1; i < len(input); i++ {
		for j := 0; j < i; j++ {
			if input[j] < input[i] && length[j]+1 >= length[i] {
				length[i] = length[j] + 1
				sequence[i] = j
			}
		}
	}

	fmt.Println(length)
	max := 0
	for i := range length {
		if length[i] > length[max] {
			max = i
		}
	}

	result := make([]int, 0)
	i := max
	for i != -1 {
		result = append(result, input[i])
		i = sequence[i]
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
