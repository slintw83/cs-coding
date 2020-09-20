package main

func MoveElementToEnd(array []int, toMove int) []int {
	end := len(array) - 1
	for i := 0; i < end; {
		if array[i] == toMove {
			array[i], array[end] = array[end], array[i]
			end--
		} else {
			i++
		}
	}

	return array
}
