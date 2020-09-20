package main

func FourNumberSum(array []int, target int) [][]int {
	result := make([][]int, 0)
	cache := make(map[int][][]int)

	for i := 0; i < len(array); i++ {
		for j := i + 1; j < len(array); j++ {
			currentSum := array[i] + array[j]
			difference := target - currentSum
			if pairs, ok := cache[difference]; ok {
				for _, pair := range pairs {
					newQuad := append(pair, array[i], array[j])
					result = append(result, newQuad)
				}
			}
		}

		for k := 0; k < i; k++ {
			currentSum := array[k] + array[i]
			cache[currentSum] = append(cache[currentSum], []int{array[k], array[i]})
		}
	}
	return result
}
