package main

func SameBSTs(arrayOne, arrayTwo []int) bool {
	if len(arrayOne) != len(arrayTwo) {
		return false
	}

	if len(arrayOne) == 0 && len(arrayTwo) == 0 {
		return true
	}

	if arrayOne[0] != arrayTwo[0] {
		return false
	}

	smallerOne := getSmaller(arrayOne)
	greaterOne := getGreaterOrEqual(arrayOne)
	smallerTwo := getSmaller(arrayTwo)
	greaterTwo := getGreaterOrEqual(arrayTwo)

	if SameBSTs(smallerOne, smallerTwo) && SameBSTs(greaterOne, greaterTwo) {
		return true
	}

	return false
}

func getSmaller(array []int) []int {
	result := []int{}

	for i := 1; i < len(array); i++ {
		if array[i] < array[0] {
			result = append(result, array[i])
		}
	}

	return result
}

func getGreaterOrEqual(array []int) []int {
	result := []int{}
	for i := 1; i < len(array); i++ {
		if array[i] >= array[0] {
			result = append(result, array[i])
		}
	}
	return result
}
