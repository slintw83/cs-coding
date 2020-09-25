package main

func MaxSubsetSumNoAdjacent(array []int) int {
	if len(array) == 0 {
		return 0
	} else if len(array) == 1 {
		return array[0]
	}

	dp := make([]int, len(array))
	dp[0] = array[0]
	dp[1] = Max(array[0], array[1])

	for i := 2; i < len(array); i++ {
		dp[i] = Max(dp[i-1], dp[i-2]+array[i])
	}

	return dp[len(array)-1]
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
