package main

func SearchInSortedMatrix(matrix [][]int, target int) []int {
	row, col := 0, len(matrix[0])-1
	for row < len(matrix) && col >= 0 {
		if matrix[row][col] == target {
			return []int{row, col}
		} else if matrix[row][col] < target {
			row++
		} else if matrix[row][col] > target {
			col--
		}
	}
	return []int{-1, -1}
}
