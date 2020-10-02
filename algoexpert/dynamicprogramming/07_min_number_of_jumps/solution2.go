package main

func MinNumberOfJumps(array []int) int {
	if len(array) == 1 {
		return 0
	}
	jumps := 0
	maxReach := array[0]
	steps := array[0]
	for i := 1; i < len(array)-1; i++ {
		maxReach = Max(maxReach, array[i]+i)
		steps--
		if steps == 0 {
			jumps++
			steps = maxReach - i
		}
	}
	return jumps + 1
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
