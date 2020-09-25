package main

func NumberOfWaysToMakeChange(n int, denoms []int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	for _, denom := range denoms {
		for amount := 1; amount <= n; amount++ {
			if denom <= amount {
				dp[amount] += dp[amount-denom]
			}
		}
	}

	return dp[n]
}
