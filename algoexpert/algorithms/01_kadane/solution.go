package main

// KadanesAlgorithm returns the maximum sum subarray
func KadanesAlgorithm(array []int) int {
	maxEndingHere := array[0]
	maxEndingSoFar := array[0]
	for i := 1; i < len(array); i++ {
		num := array[i]
		maxEndingHere = Max(num, maxEndingHere+num)
		maxEndingSoFar = Max(maxEndingHere, maxEndingSoFar)
	}
	return maxEndingSoFar
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
