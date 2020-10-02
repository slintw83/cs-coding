package main

import "fmt"

func KnapsackProblem(items [][]int, capacity int) []interface{} {
	values := make([][]int, len(items)+1)
	for i := range values {
		values[i] = make([]int, capacity+1)
	}

	for i := 1; i <= len(items); i++ {
		for j := 0; j <= capacity; j++ {
			v := items[i-1][0]
			w := items[i-1][1]
			if w <= j {
				values[i][j] = Max(values[i-1][j], values[i-1][j-w]+v)
			} else {
				values[i][j] = values[i-1][j]
			}
		}
	}

	fmt.Println(values)

	return []interface{}{
		values[len(values)-1][len(values[0])-1], // total value
		backtrack(values, items),                // item indices
	}
}

func backtrack(values [][]int, items [][]int) []int {
	result := make([]int, 0)
	i, j := len(values)-1, len(values[0])-1
	for i > 0 {
		if values[i][j] == values[i-1][j] {
			i--
		} else {
			result = append(result, i-1)
			j -= items[i-1][1]
			i--
		}
		if j == 0 {
			break
		}
	}
	reverse(result)
	return result
}

func reverse(values []int) {
	for i, j := 0, len(values)-1; i < j; i, j = i+1, j-1 {
		values[i], values[j] = values[j], values[i]
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
