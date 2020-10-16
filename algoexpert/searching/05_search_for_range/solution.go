package main

import "fmt"

func SearchForRange(array []int, target int) []int {
	start := binarySearch1(array, target)
	if start == -1 {
		return []int{-1, -1}
	}
	end := binarySearch2(array, target)
	return []int{start, end}
}

func binarySearch1(array []int, target int) int {
	start, end := 0, len(array)-1
	for start <= end {
		mid := start + (end-start)/2
		if array[mid] >= target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	fmt.Println(start, end)
	if array[start] != target {
		return -1
	}
	return start
}

func binarySearch2(array []int, target int) int {
	start, end := 0, len(array)-1
	for start <= end {
		mid := start + (end-start)/2
		if array[mid] <= target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return end
}
