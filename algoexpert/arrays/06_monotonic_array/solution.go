package main

func IsMonotonic(array []int) bool {
	if len(array) < 2 {
		return true
	}

	var increasing bool
	i := 0
	for i < len(array)-1 {
		if array[i] < array[i+1] {
			increasing = true
			break
		} else if array[i] > array[i+1] {
			increasing = false
			break
		} else {
			i++
		}
	}

	for ; i < len(array)-1; i++ {
		if (increasing && array[i] > array[i+1]) || (!increasing && array[i] < array[i+1]) {
			return false
		}
	}

	return true
}
