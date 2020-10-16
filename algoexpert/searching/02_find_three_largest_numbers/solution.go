package main

import "math"

func FindThreeLargestNumbers(array []int) []int {
	result := make([]int, 3)
	for i := range result {
		result[i] = math.MinInt32
	}
	for _, val := range array {
		if val > result[0] {
			insert(result, val)
		}
	}
	return result
}

func insert(array []int, val int) {
	array[0] = val
	if array[0] > array[1] {
		array[0], array[1] = array[1], array[0]
	}
	if array[1] > array[2] {
		array[1], array[2] = array[2], array[1]
	}
}
