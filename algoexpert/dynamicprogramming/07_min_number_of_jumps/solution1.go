package main

import "math"

// MinNumberOfJumps O(n^2)
func MinNumberOfJumps(array []int) int {
	jumps := make([]int, len(array))
	for i := 1; i < len(jumps); i++ {
		jumps[i] = math.MaxInt32
	}

	for i := 1; i < len(jumps); i++ {
		for j := 0; j < i; j++ {
			if j+array[j] >= i {
				jumps[i] = Min(jumps[i], jumps[j]+1)
			}
		}
	}

	return jumps[len(jumps)-1]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
