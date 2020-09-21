package main

import "math"

func LargestRange(array []int) []int {
	cache := make(map[int]bool)
	length, left, right := math.MinInt32, 0, 0

	for _, v := range array {
		cache[v] = true
	}

	for _, v := range array {
		if p := cache[v]; p {
			l, lv, rv := check(v, cache)
			if l > length {
				length, left, right = l, lv, rv
			}
		}
	}

	return []int{left, right}
}

func check(v int, cache map[int]bool) (int, int, int) {
	cache[v] = false
	currentLength := 1
	left, right := v-1, v+1

	for cache[left] {
		cache[left] = false
		left--
		currentLength++
	}

	for cache[right] {
		cache[right] = false
		right++
		currentLength++
	}

	return currentLength, left + 1, right - 1
}
