package main

func Powerset(array []int) [][]int {
	result := make([][]int, 0)
	result = append(result, []int{})

	for _, v := range array {
		length := len(result)
		for i := 0; i < length; i++ {
			newSubset := clone(result[i])
			newSubset = append(newSubset, v)
			result = append(result, newSubset)
		}
	}

	return result
}

func clone(arr []int) []int {
	newArr := make([]int, len(arr))
	copy(newArr, arr)
	return newArr
}
