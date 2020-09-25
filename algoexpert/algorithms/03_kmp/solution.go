package main

func KnuthMorrisPrattAlgorithm(str, substr string) bool {
	pattern := buildPattern(substr)
	return match(str, substr, pattern)
}

func match(str string, substr string, pattern []int) bool {
	i, j := 0, 0
	for i+len(substr)-j <= len(str) {
		if str[i] == substr[j] {
			if j == len(substr)-1 {
				return true
			}
			i++
			j++
		} else if j > 0 {
			j = pattern[j-1] + 1
		} else {
			i++
		}
	}

	return false
}

func buildPattern(str string) []int {
	pattern := make([]int, len(str))
	for i := range str {
		pattern[i] = -1
	}

	i, j := 1, 0
	for i < len(str) {
		if str[j] == str[i] {
			pattern[i] = j
			i++
			j++
		} else if j > 0 {
			j = pattern[j-1] + 1
		} else {
			i++
		}
	}
	return pattern
}
