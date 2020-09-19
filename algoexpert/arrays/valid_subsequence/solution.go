package valid_subsequence

func IsValidSubsequence(array []int, sequence []int) bool {
	i, j := 0, 0

	for i < len(array) {
		if array[i] == sequence[j] {
			j++
			if j >= len(sequence) {
				return true
			}
		}
		i++
	}

	return false
}
