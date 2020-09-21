package main

import "fmt"

/// MinRewards naives solution O(n^2)
func MinRewards(scores []int) int {
	rewards := make([]int, len(scores))
	rewards[0] = 1
	for i := 1; i < len(scores); i++ {
		if scores[i] < scores[i-1] {
			rewards[i] = 1
			fix(scores, rewards, i-1)
		} else {
			rewards[i] = rewards[i-1] + 1
		}
	}

	fmt.Println(rewards) // Degub

	sum := 0
	for _, v := range rewards {
		sum += v
	}
	return sum
}

func fix(scores []int, rewards []int, start int) {
	for start >= 0 {
		if rewards[start] < rewards[start+1] {
			return
		}
		rewards[start] = max(rewards[start], rewards[start+1]+1)
		start--
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
