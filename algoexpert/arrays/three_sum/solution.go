package three_sum

import "sort"

func ThreeNumberSum(array []int, target int) [][]int {
	result := make([][]int, 0)
	sort.Ints(array)
	for i, v := range array {
		l, r := i+1, len(array)-1

		for l < r {
			current_sum := v + array[l] + array[r]
			if current_sum == target {
				result = append(result, []int{v, array[l], array[r]})
			}

			if current_sum < target {
				l++
			} else {
				r--
			}
		}
	}
	return result
}
