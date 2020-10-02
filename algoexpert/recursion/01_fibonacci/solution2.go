package main

func GetNthFib(n int) int {
	return getNthFib(n, make(map[int]int))
}

func getNthFib(n int, dp map[int]int) int {
	if v, ok := dp[n]; ok {
		return v
	} else if n == 1 {
		dp[n] = 0
	} else if n == 2 {
		dp[n] = 1
	} else {
		dp[n] = getNthFib(n-2, dp) + getNthFib(n-1, dp)
	}
	return dp[n]
}
