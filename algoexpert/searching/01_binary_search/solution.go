package main

import "fmt"

func BinarySearch(array []int, target int) int {
	start, end := 0, len(array)-1
	for start <= end {
		mid := start + (end-start)/2
		fmt.Println(array[mid])
		if array[mid] == target {
			return mid
		}
		if target > array[mid] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
