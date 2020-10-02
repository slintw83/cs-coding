package main

import "fmt"

// O(nm*min(n, m)) time | O(nm*min(n, m)) space
func LongestCommonSubsequence(s1 string, s2 string) string {
	lcs := make([][]int, len(s2)+1)
	for i := range lcs {
		lcs[i] = make([]int, len(s1)+1)
	}

	for i := 1; i <= len(s2); i++ {
		for j := 1; j <= len(s1); j++ {
			if s2[i-1] == s1[j-1] {
				lcs[i][j] = lcs[i-1][j-1] + 1
			} else {
				lcs[i][j] = max(lcs[i-1][j], lcs[i][j-1])
			}
		}
	}

	fmt.Println(lcs)
	return build(lcs, s1)
}

func build(lcs [][]int, str string) string {
	i := len(lcs) - 1
	j := len(lcs[0]) - 1
	sequence := make([]byte, 0)

	for i != 0 && j != 0 {
		if lcs[i][j] == lcs[i-1][j] {
			i--
		} else if lcs[i][j] == lcs[i][j-1] {
			j--
		} else {
			sequence = append(sequence, str[j-1])
			i--
			j--
		}
	}

	for i, j := 0, len(sequence)-1; i < j; i, j = i+1, j-1 {
		sequence[i], sequence[j] = sequence[j], sequence[i]
	}
	return string(sequence)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
