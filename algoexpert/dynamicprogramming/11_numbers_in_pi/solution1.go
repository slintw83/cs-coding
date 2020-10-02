package main

import "math"

func NumbersInPi(pi string, numbers []string) int {
	cache := make(map[int]int)
	numbersMap := make(map[string]bool)
	for _, num := range numbers {
		numbersMap[num] = true
	}
	minSpaces := recurse(pi, numbersMap, cache, 0)
	if minSpaces == math.MaxInt32 {
		return -1
	} else {
		return minSpaces
	}
}

func recurse(pi string, numbers map[string]bool, cache map[int]int, idx int) int {
	if idx == len(pi) {
		return -1
	} else if val, ok := cache[idx]; ok {
		return val
	}

	minSpaces := math.MaxInt32
	for i := idx; i < len(pi); i++ {
		prefix := pi[idx : i+1]
		if numbers[prefix] {
			minSpacesSuffix := recurse(pi, numbers, cache, i+1)
			minSpaces = Min(minSpaces, minSpacesSuffix+1)
		}
	}
	cache[idx] = minSpaces
	return cache[idx]
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
