package main

import "fmt"

/// MinRewards using local minimums
/// O(n)
func MinRewards(scores []int) int {
	minIdx := findLocalMins(scores)
	rewards := make([]int, len(scores))
	for i := range rewards {
		rewards[i] = 1
	}

	fmt.Println(minIdx)

	for _, idx := range minIdx {
		expand(scores, rewards, idx)
	}

	fmt.Println(rewards)

	sum := 0
	for _, v := range rewards {
		sum += v
	}
	return sum
}

func expand(scores []int, rewards []int, idx int) {
	leftIdx := idx - 1
	for leftIdx >= 0 && scores[leftIdx] > scores[leftIdx+1] {
		rewards[leftIdx] = max(rewards[leftIdx], rewards[leftIdx+1]+1)
		leftIdx--
	}
	rightIdx := idx + 1
	for rightIdx < len(rewards) && scores[rightIdx] > scores[rightIdx-1] {
		rewards[rightIdx] = rewards[rightIdx-1] + 1
		rightIdx++
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findLocalMins(scores []int) []int {
	localMins := make([]int, 0)
	if len(scores) == 1 {
		localMins = append(localMins, 0)
		return localMins
	}

	for i, v := range scores {
		if i == 0 && v < scores[i+1] {
			localMins = append(localMins, i)
		}
		if i == len(scores)-1 && v < scores[i-1] {
			localMins = append(localMins, i)
		}
		if i == 0 || i == len(scores)-1 {
			continue
		}

		if v < scores[i+1] && v < scores[i-1] {
			localMins = append(localMins, i)
		}
	}

	return localMins
}
