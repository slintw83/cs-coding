package main

import "math"

func MaxProfitWithKTransactions(prices []int, k int) int {
	if len(prices) == 0 {
		return 0
	}

	profit := make([][]int, k+1)
	for i := range profit {
		profit[i] = make([]int, len(prices))
	}

	for t := 1; t < k+1; t++ {
		maxThusFar := math.MinInt32
		for d := 1; d < len(prices); d++ {
			maxThusFar = Max(maxThusFar, profit[t-1][d-1]-prices[d-1]) // maxium profit (after buying) before d
			profit[t][d] = Max(profit[t][d-1], maxThusFar+prices[d])
		}
	}

	return profit[k][len(prices)-1]
}

func Max(val int, rest ...int) int {
	curr := val
	for _, num := range rest {
		if num > curr {
			curr = num
		}
	}
	return curr
}
