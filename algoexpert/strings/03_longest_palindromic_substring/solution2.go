package main

// LongestPalindromicSubstring O(n^2) time and O(1) space
func LongestPalindromicSubstring(str string) string {
	if len(str) <= 1 {
		return str
	}

	start, maxlength := 1, 1
	for i := range str {
		leftidx := i - 1
		rightidx := i
		for leftidx >= 0 && rightidx < len(str) && str[leftidx] == str[rightidx] {
			if rightidx-leftidx+1 > maxlength {
				maxlength = rightidx - leftidx + 1
				start = leftidx
			}
			leftidx--
			rightidx++
		}

		leftidx = i - 1
		rightidx = i + 1
		for leftidx >= 0 && rightidx < len(str) && str[leftidx] == str[rightidx] {
			if rightidx-leftidx+1 > maxlength {
				maxlength = rightidx - leftidx + 1
				start = leftidx
			}
			leftidx--
			rightidx++
		}
	}

	return str[start : start+maxlength]
}
