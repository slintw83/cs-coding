package main

import "math"

func MinNumberOfCoinsForChange(n int, denoms []int) int {
	numOfCoins := make([]int, n+1)

	for i := range numOfCoins {
		numOfCoins[i] = math.MaxInt32
	}
	numOfCoins[0] = 0

	for _, denom := range denoms {
		for i := range numOfCoins {
			if denom <= i {
				numOfCoins[i] = Min(numOfCoins[i], numOfCoins[i-denom]+1)
			}
		}
	}

	if numOfCoins[n] != math.MaxInt32 {
		return numOfCoins[n]
	}

	return -1
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
