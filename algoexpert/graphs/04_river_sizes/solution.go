package main

func RiverSizes(matrix [][]int) []int {
	visited := make([][]bool, len(matrix))
	sizes := make([]int, 0)
	for i := range visited {
		visited[i] = make([]bool, len(matrix[0]))
	}

	for row := range matrix {
		for col := range matrix[row] {
			if !visited[row][col] && matrix[row][col] == 1 {
				size := dfs(matrix, visited, row, col)
				sizes = append(sizes, size)
			}
		}
	}

	return sizes
}

func dfs(matrix [][]int, visited [][]bool, row, col int) int {
	if row < 0 || row >= len(matrix) {
		return 0
	} else if col < 0 || col >= len(matrix[0]) {
		return 0
	} else if matrix[row][col] == 0 || visited[row][col] {
		return 0
	}
	visited[row][col] = true
	size := 1
	size += dfs(matrix, visited, row-1, col)
	size += dfs(matrix, visited, row, col-1)
	size += dfs(matrix, visited, row+1, col)
	size += dfs(matrix, visited, row, col+1)
	return size
}
