package main

///LongestPalindromicSubstring O(n^2) time and space -> DP
func LongestPalindromicSubstring(str string) string {
	n := len(str)

	if n <= 1 {
		return str
	}

	if n == 2 {
		if str[0] == str[1] {
			return str
		} else {
			return str[:1]
		}
	}

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	start, maxlength := 0, 0

	// Palindromes of size 1
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}

	// Palindromes of size 2
	for i := 0; i < n-1; i++ {
		if str[i] == str[i+1] {
			dp[i][i+1] = true
			start = i
			maxlength = 2
		}
	}

	for k := 3; k <= n; k++ {
		for i := 0; i < (n - k + 1); i++ {
			j := i + k - 1 //Ending index

			if dp[i+1][j-1] && str[i] == str[j] {
				dp[i][j] = true
				if k > maxlength {
					maxlength = k
					start = i
				}
			}
		}
	}

	return str[start : start+maxlength]
}
