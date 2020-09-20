package main

import (
	"math"
	"sort"
)

func SmallestDifference(array1, array2 []int) []int {
	sort.Ints(array1)
	sort.Ints(array2)

	i, j := 0, 0
	min_dif := math.MaxInt64
	var result []int

	for {
		dif := Abs(array1[i] - array2[j])
		if dif < min_dif {
			min_dif = dif
			result = []int{array1[i], array2[j]}

			if dif == 0 {
				break
			}
		}

		if array1[i] < array2[j] {
			i++
		} else {
			j++
		}

		if i == len(array1) || j == len(array2) {
			break
		}
	}

	return result
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	SmallestDifference([]int{-1, 5, 10, 20, 28, 3}, []int{26, 134, 135, 15, 17})
}
