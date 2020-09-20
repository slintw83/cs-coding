package main

func SpiralTraverse(array [][]int) []int {
	result := make([]int, 0)

	startRow, endRow := 0, len(array)-1
	startCol, endCol := 0, len(array[0])-1

	for startRow <= endRow && startCol <= endCol {
		for col := startCol; col <= endCol; col++ {
			result = append(result, array[startRow][col])
		}

		for row := startRow + 1; row <= endRow; row++ {
			result = append(result, array[row][endCol])
		}

		for col := endCol - 1; col >= startCol; col-- {
			if startRow == endRow {
				break
			}
			result = append(result, array[endRow][col])
		}

		for row := endRow - 1; row > startRow; row-- {
			if startCol == endCol {
				break
			}
			result = append(result, array[row][startCol])
		}

		startRow++
		endRow--
		startCol++
		endCol--
	}

	return result
}
