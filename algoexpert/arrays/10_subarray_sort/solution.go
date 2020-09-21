package main

import "math"

/// Find the min and max unsorted values
/// Find thier correct position and return
func SubarraySort(array []int) []int {
	minUnsorted, maxUnsorted := math.MaxInt32, math.MinInt32
	for i, v := range array {
		if isOutOfOrder(i, v, array) {
			if v < minUnsorted {
				minUnsorted = v
			}

			if v > maxUnsorted {
				maxUnsorted = v
			}
		}
	}

	if minUnsorted == math.MaxInt32 {
		return []int{-1, -1}
	}

	leftIdx := 0
	for minUnsorted >= array[leftIdx] {
		leftIdx++
	}

	rightIdx := len(array) - 1
	for maxUnsorted <= array[rightIdx] {
		rightIdx--
	}

	return []int{leftIdx, rightIdx}
}

func isOutOfOrder(i int, v int, array []int) bool {
	if i == 0 {
		return v > array[i+1]
	} else if i == len(array)-1 {
		return v < array[i-1]
	}
	return v < array[i-1] || v > array[i+1]
}
