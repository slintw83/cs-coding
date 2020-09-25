package main

func LevenshteinDistance(a, b string) int {
	dp := make([][]int, len(b)+1)

	for i := range dp {
		dp[i] = make([]int, len(a)+1)
	}

	for i := range dp[0] {
		dp[0][i] = i
	}

	for i := range dp {
		dp[i][0] = i
	}

	for i := 1; i < len(b)+1; i++ {
		for j := 1; j < len(a)+1; j++ {
			if b[i-1] == a[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + Min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}

	return dp[len(b)][len(a)]
}

func Min(args ...int) int {
	curr := args[0]
	for i := range args {
		if curr > args[i] {
			curr = args[i]
		}
	}
	return curr
}
