package main

func WaterArea(heights []int) int {
	leftTallest := make([]int, len(heights))
	rightTallest := make([]int, len(heights))
	areas := make([]int, len(heights))

	max := 0
	for i := range heights {
		leftTallest[i] = max
		max = Max(max, heights[i])
	}

	max = 0
	for i := len(heights) - 1; i >= 0; i-- {
		rightTallest[i] = max
		max = Max(max, heights[i])
	}

	for i := range heights {
		min := Min(leftTallest[i], rightTallest[i])
		if heights[i] < min {
			areas[i] = min - heights[i]
		}
	}

	sum := 0
	for i := range areas {
		sum += areas[i]
	}

	return sum
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
