package main

func MinRewards(scores []int) int {
	rewards := make([]int, len(scores))
	for i := range rewards {
		rewards[i] = 1
	}

	for i := 1; i < len(scores); i++ {
		if scores[i] > scores[i-1] {
			rewards[i] = rewards[i-1] + 1
		}
	}

	for i := len(scores) - 2; i >= 0; i-- {
		if scores[i] > scores[i+1] {
			rewards[i] = max(rewards[i], rewards[i+1]+1)
		}
	}

	sum := 0
	for _, v := range rewards {
		sum += v
	}
	return sum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
