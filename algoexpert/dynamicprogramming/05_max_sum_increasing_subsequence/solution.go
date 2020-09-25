package main

import "fmt"

func MaxSumIncreasingSubsequence(array []int) []interface{} {
	sums := make([]int, len(array))
	sums[0] = array[0]

	sequences := make([]int, len(array))
	for i := range sequences {
		sums[i] = array[i]
		sequences[i] = -1
	}

	for i := 1; i < len(array); i++ {
		for k := 0; k < i; k++ {
			if array[k] < array[i] {
				if sums[k]+array[i] > sums[i] {
					sums[i] = sums[k] + array[i]
					sequences[i] = k
				}
			}
		}
	}

	fmt.Println(sums)

	maxidx := 0
	for i := range sums {
		if sums[i] > sums[maxidx] {
			maxidx = i
		}
	}

	result := make([]int, 0)
	for i := maxidx; i != -1; i = sequences[i] {
		result = append([]int{array[i]}, result...)
	}

	// Write your code here.
	return []interface{}{
		sums[maxidx], // Example max sum
		result,       // Example max sequence
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
